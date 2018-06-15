# REST vs gRPC benchmark

### dependencies:
- Go 1.9+
- protoc 3.0+

### build:
`make build`

### run benchmarks:
- `./rest_vs_grpc_benchmark`
- `cd rest_client`
- `go test -bench=.`

to compare with pure gRPC:
- `cd rpc_client`
- `go test -bench=.`
