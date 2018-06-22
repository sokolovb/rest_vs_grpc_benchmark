package rpc_client

import (
	"context"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/proto"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/rest_vs_grpc"
	"google.golang.org/grpc"
	"io"
)

type RpcClient struct {
	c BenchmarkClient
}

func NewRpcClient() *RpcClient {
	conn, _ := grpc.Dial("localhost:"+PortRpc, grpc.WithInsecure())
	return &RpcClient{NewBenchmarkClient(conn)}
}

func (rpc *RpcClient) GetInt() error {
	_, err := rpc.c.GetInt(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetIntSlice() error {
	_, err := rpc.c.GetIntSlice(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetString() error {
	_, err := rpc.c.GetString(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetStringSlice() error {
	_, err := rpc.c.GetStringSlice(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetBlob() error {
	_, err := rpc.c.GetBlob(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetStruct() error {
	_, err := rpc.c.GetStruct(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetStructSlices() error {
	_, err := rpc.c.GetStructSlices(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetStructStructs() error {
	_, err := rpc.c.GetStructStructs(context.Background(), &Request{})
	return err
}

func (rpc *RpcClient) GetFileStream() error {
	stream, err := rpc.c.GetFileStream(context.Background(), &Request{})
	if err != nil {
		return err
	}
	for {
		_, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				return stream.CloseSend()
			}
		}
	}
}