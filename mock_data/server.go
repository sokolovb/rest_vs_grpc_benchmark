package mock_data

import (
	. "github.com/sokolovb/rest_vs_grpc_benchmark/proto"
	"math/rand"
	"reflect"
)

const (
	sliceSize = 1000
	blobSize  = 4194000
	str       = "qwerty"
)

type Data struct {
	integer          *Int
	integerSlice     *IntSlice
	str              *String
	strSlice         *StringSlice
	blob             *Blob
	structure        *Struct
	structureSlices  *StructSlices
	structureStructs *StructStructs
}

func NewData() *Data {
	s := &Data{
		integer:          new(Int),
		integerSlice:     new(IntSlice),
		str:              new(String),
		strSlice:         new(StringSlice),
		blob:             new(Blob),
		structure:        new(Struct),
		structureSlices:  new(StructSlices),
		structureStructs: new(StructStructs),
	}
	s.integer.Value = rand.Int31()
	s.integerSlice.Value = make([]int32, sliceSize)
	s.str.Value = str
	s.strSlice.Value = make([]string, sliceSize)
	s.blob.Value = make([]byte, blobSize)
	for i := 0; i < sliceSize; i++ {
		s.integerSlice.Value[i] = rand.Int31()
		s.strSlice.Value[i] = str
	}
	for i := 0; i < blobSize; i++ {
		s.blob.Value[i] = byte(rand.Intn(255))
	}

	vPtr := reflect.ValueOf(s.structure)
	v := reflect.ValueOf(*s.structure)
	for i := 0; i < v.NumField(); i++ {
		vPtr.Elem().Field(i).SetInt(rand.Int63())
	}

	vPtr = reflect.ValueOf(s.structureSlices)
	v = reflect.ValueOf(*s.structureSlices)
	for i := 0; i < v.NumField(); i++ {
		vPtr.Elem().Field(i).Set(reflect.ValueOf(s.integerSlice))
	}

	vPtr = reflect.ValueOf(s.structureStructs)
	v = reflect.ValueOf(*s.structureStructs)
	for i := 0; i < v.NumField(); i++ {
		vPtr.Elem().Field(i).Set(reflect.ValueOf(s.structure))
	}
	return s
}

func (s *Data) GetInt() *Int {
	return s.integer
}

func (s *Data) GetIntSlice() *IntSlice {
	return s.integerSlice
}

func (s *Data) GetString() *String {
	return s.str
}

func (s *Data) GetStringSlice() *StringSlice {
	return s.strSlice
}

func (s *Data) GetBlob() *Blob {
	return s.blob
}

func (s *Data) GetStruct() *Struct {
	return s.structure
}

func (s *Data) GetStructSlices() *StructSlices {
	return s.structureSlices
}

func (s *Data) GetStructStructs() *StructStructs {
	return s.structureStructs
}
