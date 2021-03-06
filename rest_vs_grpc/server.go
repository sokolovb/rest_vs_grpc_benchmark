package rest_vs_grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	md "github.com/sokolovb/rest_vs_grpc_benchmark/mock_data"
	"github.com/sokolovb/rest_vs_grpc_benchmark/proto"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"io"
)

const (
	PortRestProxy = "8080"
	PortRpc       = "50030"
)

func ServeRPC(data *md.Data, port string) error {
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return fmt.Errorf("failed to initialize TCP listen: %v", err)
	}

	defer lis.Close()
	serverRpc := grpc.NewServer()
	impl := NewServer(data)

	RegisterBenchmarkServer(serverRpc, impl)
	reflection.Register(serverRpc)
	return serverRpc.Serve(lis)
}

func StartRestProxy(port string) error {
	gwmux := runtime.NewServeMux()
	err := proto.RegisterBenchmarkHandlerFromEndpoint(context.Background(), gwmux, "localhost:"+PortRpc, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return fmt.Errorf("failed to register gRPC server: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	return http.ListenAndServe("localhost:"+port, mux)
}

type Handler struct {
	data *md.Data
}

func NewServer(data *md.Data) BenchmarkServer {
	return &Handler{data}
}

func (h *Handler) GetInt(context.Context, *Request) (*Int, error) {
	return h.data.GetInt(), nil
}

func (h *Handler) GetIntSlice(context.Context, *Request) (*IntSlice, error) {
	return h.data.GetIntSlice(), nil
}

func (h *Handler) GetString(context.Context, *Request) (*String, error) {
	return h.data.GetString(), nil
}

func (h *Handler) GetStringSlice(context.Context, *Request) (*StringSlice, error) {
	return h.data.GetStringSlice(), nil
}

func (h *Handler) GetBlob(context.Context, *Request) (*Blob, error) {
	return h.data.GetBlob(), nil
}

func (h *Handler) GetStruct(context.Context, *Request) (*Struct, error) {
	return h.data.GetStruct(), nil
}

func (h *Handler) GetStructSlices(context.Context, *Request) (*StructSlices, error) {
	return h.data.GetStructSlices(), nil
}

func (h *Handler) GetStructStructs(context.Context, *Request) (*StructStructs, error) {
	return h.data.GetStructStructs(), nil
}

func (h *Handler) GetFileStream(req *Request, stream Benchmark_GetFileStreamServer) error {
	buf := make([]byte, 32 * 1024)
	file := h.data.FileReader()
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				return nil
			}
		}
		if err := stream.Send(&Blob{Value: buf[:n]}); err != nil {
			return err
		}
	}
	return nil
}