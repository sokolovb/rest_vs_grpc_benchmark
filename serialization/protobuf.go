package serialization

import "github.com/golang/protobuf/proto"

func SerializeProtobuf(in proto.Message) error {
	_, err := proto.Marshal(in)
	return err
}

func DeserializeProtobuf(in []byte, out proto.Message) error {
	return proto.Unmarshal(in, out)
}