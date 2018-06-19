package serialization

import "github.com/golang/protobuf/proto"

func SerializeProtobuf(in proto.Message) error {
	_, err := proto.Marshal(in)
	return err
}

func DeserializeProtobuf(in []byte) error {
	var out proto.Message
	return proto.Unmarshal(in, out)
}