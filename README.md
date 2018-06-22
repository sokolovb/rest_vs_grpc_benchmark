# REST vs gRPC benchmark

### dependencies:
- Go 1.9+
- protoc 3.0+

### build:
`make build`

### run benchmarks:
- `./rest_vs_grpc_benchmark`
- `cd rest_client`
- `sh run_bench.sh `

to compare with pure gRPC:
- `cd rpc_client`
- `sh run_bench.sh `

serialization:
- `cd serialization`
- `sh run_bench.sh `