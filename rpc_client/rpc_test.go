package rpc_client_test

import (
	"os"
	"rest-vs-grpc/rpc_client"
	"testing"
)

var rpc *rpc_client.RpcClient

func TestMain(m *testing.M) {
	rpc = rpc_client.NewRpcClient()
	os.Exit(m.Run())
}

func BenchmarkRpcClient_GetInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetInt(); err != nil {
			b.Fatalf("GetInt() error: %v", err)
		}
	}
}

func BenchmarkRpcClient_GetBlob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
		}
	}
}
