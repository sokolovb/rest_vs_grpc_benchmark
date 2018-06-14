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

func BenchmarkRpcClient_GetIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetIntSlice(); err != nil {
			b.Fatalf("GetIntSlice() error: %v", err)
		}
	}
}

func BenchmarkRpcClient_GetString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetString(); err != nil {
			b.Fatalf("GetString() error: %v", err)
		}
	}
}

func BenchmarkRpcClient_GetStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetStringSlice(); err != nil {
			b.Fatalf("GetStringSlice() error: %v", err)
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

func BenchmarkRpcClient_GetStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := rpc.GetStruct(); err != nil {
			b.Fatalf("GetStruct() error: %v", err)
		}
	}
}