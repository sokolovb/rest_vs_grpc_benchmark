package main

import (
	. "github.com/sokolovb/rest_vs_grpc_benchmark/mock_data"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/pure_rest"
	. "github.com/sokolovb/rest_vs_grpc_benchmark/rest_vs_grpc"
	"log"
)

//go:generate protoc -I/usr/local/include -I. -I$GOPATH/src -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. proto/api.proto

func main() {
	data := NewData()
	log.Print("Data created")

	go func() {
		log.Printf("Pure REST server started on port %s", PortRest)
		log.Fatal(StartPureRestServer(data, PortRest))
	}()

	go func() {
		log.Printf("gRPC server started on port %s", PortRpc)
		log.Fatal(ServeRPC(data, PortRpc))
	}()

	log.Printf("Proxy REST server started on port %s", PortRestProxy)
	log.Fatal(StartRestProxy(PortRestProxy))
}
