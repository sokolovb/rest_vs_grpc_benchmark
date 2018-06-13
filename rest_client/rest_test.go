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

func BenchmarkRestClient_GetBlob_Proxy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restProxy.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
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

func BenchmarkRestClient_GetBlob_Pure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := restPure.GetBlob(); err != nil {
			b.Fatalf("GetBlob() error: %v", err)
		}
	}
}
