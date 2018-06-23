# REST vs gRPC benchmark

### dependencies:
- Go 1.9+
- protoc 3.0+
- plotly (https://plot.ly/python/) with credentials

### build:
`make build`

### run benchmarks:
- `./rest_vs_grpc_benchmark`

keep the server listening

- `sh run_bench.sh ${plotly_username} ${plotly_token}`
