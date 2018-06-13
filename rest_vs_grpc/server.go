package rest_vs_grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	md "rest-vs-grpc/mock_data"
	. "rest-vs-grpc/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	"rest-vs-grpc/proto"
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
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(assetFS())))
	return http.ListenAndServe("localhost:"+port, mux)
}

type Handler struct {
	data *md.Data
}

func NewServer(data *md.Data) BenchmarkServer {
	return &Handler{data}
}

func (h *Handler) GetInt(context.Context, *Request) (*Int, error) {
	return &Int{h.data.GetInt()}, nil
}

func (h *Handler) GetIntSlice(context.Context, *Request) (*IntSlice, error) {
	return &IntSlice{h.data.GetIntSlice()}, nil
}

func (h *Handler) GetString(context.Context, *Request) (*String, error) {
	return &String{h.data.GetString()}, nil
}

func (h *Handler) GetStringSlice(context.Context, *Request) (*StringSlice, error) {
	return &StringSlice{h.data.GetStringSlice()}, nil
}

func (h *Handler) GetBlob(context.Context, *Request) (*Blob, error) {
	return &Blob{h.data.GetBlob()}, nil
}

func (h *Handler) GetStruct(context.Context, *Request) (*Struct, error) {
	return h.data.GetStruct(), nil
}
