package serialization_test

import (
	"github.com/gin-gonic/gin/json"
	pb "github.com/golang/protobuf/proto"
	"github.com/sokolovb/rest_vs_grpc_benchmark/mock_data"
	"github.com/sokolovb/rest_vs_grpc_benchmark/proto"
	"github.com/sokolovb/rest_vs_grpc_benchmark/serialization"
	"os"
	"testing"
)

var data *mock_data.Data

var serializedJSON, serializedProtobuf struct {
	integer          []byte
	integerSlice     []byte
	str              []byte
	strSlice         []byte
	blob             []byte
	structure        []byte
	structureSlices  []byte
	structureStructs []byte
}

func TestMain(m *testing.M) {
	data = mock_data.NewData()

	serializedJSON.integer, _ = json.Marshal(data.GetInt())
	serializedJSON.integerSlice, _ = json.Marshal(data.GetIntSlice())
	serializedJSON.str, _ = json.Marshal(data.GetString())
	serializedJSON.strSlice, _ = json.Marshal(data.GetStringSlice())
	serializedJSON.blob, _ = json.Marshal(data.GetBlob())
	serializedJSON.structure, _ = json.Marshal(data.GetStruct())
	serializedJSON.structureSlices, _ = json.Marshal(data.GetStructSlices())
	serializedJSON.structureStructs, _ = json.Marshal(data.GetStructStructs())

	serializedProtobuf.integer, _ = pb.Marshal(data.GetInt())
	serializedProtobuf.integerSlice, _ = pb.Marshal(data.GetIntSlice())
	serializedProtobuf.str, _ = pb.Marshal(data.GetString())
	serializedProtobuf.strSlice, _ = pb.Marshal(data.GetStringSlice())
	serializedProtobuf.blob, _ = pb.Marshal(data.GetBlob())
	serializedProtobuf.structure, _ = pb.Marshal(data.GetStruct())
	serializedProtobuf.structureSlices, _ = pb.Marshal(data.GetStructSlices())
	serializedProtobuf.structureStructs, _ = pb.Marshal(data.GetStructStructs())

	os.Exit(m.Run())
}

// serialize json

func BenchmarkSerializeJSON_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetInt()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_intSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetIntSlice()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetString()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_stringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetStringSlice()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetStruct()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_blob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetBlob()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_structSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetStructSlices()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

func BenchmarkSerializeJSON_structStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeJSON(data.GetStructStructs()); err != nil {
			b.Fatalf("SerializeJSON() error: %v", err)
		}
	}
}

// serialize protobuf

func BenchmarkSerializeProtobuf_int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetInt()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_intSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetIntSlice()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetString()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_stringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetStringSlice()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetStruct()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_blob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetBlob()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_structSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetStructSlices()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkSerializeProtobuf_structStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := serialization.SerializeProtobuf(data.GetStructStructs()); err != nil {
			b.Fatalf("SerializeProtobuf() error: %v", err)
		}
	}
}

// deserialize json

func BenchmarkDeserializeJSON_int(b *testing.B) {
	out := new(proto.Int)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.integer, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_intSlice(b *testing.B) {
	out := new(proto.IntSlice)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.integerSlice, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_string(b *testing.B) {
	out := new(proto.String)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.str, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_stringSlice(b *testing.B) {
	out := new(proto.StringSlice)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.strSlice, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_struct(b *testing.B) {
	out := new(proto.Struct)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.structure, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_blob(b *testing.B) {
	out := new(proto.Blob)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.blob, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_structSlices(b *testing.B) {
	out := new(proto.StructSlices)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.structureSlices, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

func BenchmarkDeserializeJSON_structStructs(b *testing.B) {
	out := new(proto.StructStructs)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeJSON(serializedJSON.structureStructs, out); err != nil {
			b.Fatalf("DeserializeJSON() error: %v", err)
		}
	}
}

// deserialize protobuf

func BenchmarkDeserializeProtobuf_int(b *testing.B) {
	out := new(proto.Int)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.integer, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_intSlice(b *testing.B) {
	out := new(proto.IntSlice)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.integerSlice, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_string(b *testing.B) {
	out := new(proto.String)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.str, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_stringSlice(b *testing.B) {
	out := new(proto.StringSlice)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.strSlice, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_struct(b *testing.B) {
	out := new(proto.Struct)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.structure, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_blob(b *testing.B) {
	out := new(proto.Blob)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.blob, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_structSlices(b *testing.B) {
	out := new(proto.StructSlices)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.structureSlices, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}

func BenchmarkDeserializeProtobuf_structStructs(b *testing.B) {
	out := new(proto.StructStructs)
	for i := 0; i < b.N; i++ {
		if err := serialization.DeserializeProtobuf(serializedProtobuf.structureStructs, out); err != nil {
			b.Fatalf("DeserializeProtobuf() error: %v", err)
		}
	}
}
