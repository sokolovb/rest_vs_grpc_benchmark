package mock_data

import (
	"math/rand"
	"reflect"
	. "rest-vs-grpc/proto"
)

const (
	sliceSize = 1000
	blobSize  = 4194000
	str       = "qwerty"
)

type Data struct {
	integer      int32
	integerSlice []int32
	str          string
	strSlice     []string
	blob         []byte
	structure    *Struct
}

func NewData() *Data {
	s := &Data{
		integer:      rand.Int31(),
		integerSlice: make([]int32, sliceSize),
		str:          str,
		strSlice:     make([]string, sliceSize),
		blob:         make([]byte, blobSize),
		structure:    new(Struct),
	}
	for i := 0; i < sliceSize; i++ {
		s.integerSlice[i] = rand.Int31()
		s.strSlice[i] = str
	}
	for i := 0; i < blobSize; i++ {
		s.blob[i] = byte(rand.Intn(255))
	}
	vPtr := reflect.ValueOf(s.structure)
	v := reflect.ValueOf(*s.structure)
	for i := 0; i < v.NumField(); i++ {
		vPtr.Elem().Field(i).SetInt(rand.Int63())
	}
	return s
}

func (s *Data) GetInt() int32 {
	return s.integer
}

func (s *Data) GetIntSlice() []int32 {
	return s.integerSlice
}

func (s *Data) GetString() string {
	return s.str
}

func (s *Data) GetStringSlice() []string {
	return s.strSlice
}

func (s *Data) GetBlob() []byte {
	return s.blob
}

func (s *Data) GetStruct() *Struct {
	return s.structure
}
