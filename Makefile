build: $(GOPATH)/bin/protoc-gen-grpc-gateway $(GOPATH)/bin/protoc-gen-swagger $(GOPATH)/bin/protoc-gen-go
	go generate
	go build

$(GOPATH)/bin/protoc-gen-grpc-gateway:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

$(GOPATH)/bin/protoc-gen-swagger:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

$(GOPATH)/bin/protoc-gen-go:
	go get -u github.com/golang/protobuf/protoc-gen-go
