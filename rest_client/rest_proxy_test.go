package rest_client_test

import (
	. "github.com/sokolovb/rest_vs_grpc_benchmark/pure_rest"
	"github.com/sokolovb/rest_vs_grpc_benchmark/rest_client"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/rest_vs_grpc"
	"os"
	"testing"
)

var restProxy, restPure *rest_client.RestClient

func TestMain(m *testing.M) {
	restPure = rest_client.NewRestClient(PortRest)
	restProxy = rest_client.NewRestClient(PortRestProxy)
	os.Exit(m.Run())
}

// grpc-gateway proxy benchmarks

func BenchmarkRestClient_GetInt_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetInt(); err != nil {
			b.Fatalf("GetInt() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetIntSlice_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetIntSlice(); err != nil {
			b.Fatalf("GetIntSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetString_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetString(); err != nil {
			b.Fatalf("GetString() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStringSlice_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetStringSlice(); err != nil {
			b.Fatalf("GetStringSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStruct_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetStruct(); err != nil {
			b.Fatalf("GetStruct() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetBlob_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStructSlices_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetStructSlices(); err != nil {
			b.Fatalf("GetStructSlices() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStructStructs_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetStructStructs(); err != nil {
			b.Fatalf("GetStructStructs() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetFile_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetFile(); err != nil {
			b.Fatalf("GetFile() error: %v", err)
		}
	}
}
