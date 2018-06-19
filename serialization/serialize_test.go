package serialization_test

import (
	"testing"
	"github.com/sokolovb/rest_vs_grpc_benchmark/serialization"
	"os"
	"github.com/sokolovb/rest_vs_grpc_benchmark/mock_data"
)

var data *mock_data.Data

func TestMain(m *testing.M) {
	data = mock_data.NewData()
	os.Exit(m.Run())
}

func BenchmarkSerializeJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetBlob()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetBlob()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}