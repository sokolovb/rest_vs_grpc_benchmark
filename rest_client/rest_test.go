package rest_client_test

import (
	"os"
	. "rest-vs-grpc/pure_rest"
	"rest-vs-grpc/rest_client"
	. "rest-vs-grpc/rest_vs_grpc"
	"testing"
)

var restProxy, restPure *rest_client.RestClient

func TestMain(m *testing.M) {
	restPure = rest_client.NewRestClient(PortRest)
	restProxy = rest_client.NewRestClient(PortRestProxy)
	os.Exit(m.Run())
}

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

func BenchmarkRestClient_GetBlob_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
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

func BenchmarkRestClient_GetInt_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetInt(); err != nil {
			b.Fatalf("GetInt() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetIntSlice_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetIntSlice(); err != nil {
			b.Fatalf("GetIntSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetString_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetString(); err != nil {
			b.Fatalf("GetString() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStringSlice_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStringSlice(); err != nil {
			b.Fatalf("GetStringSlice() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetBlob_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
		}
	}
}

func BenchmarkRestClient_GetStruct_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetStruct(); err != nil {
			b.Fatalf("GetStruct() error: %v", err)
		}
	}
}