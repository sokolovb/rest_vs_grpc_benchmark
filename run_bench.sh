#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    echo "Please specify Plotly credentials"
    exit 0
fi

echo "run rpc benchmarks"
cd rpc_client
sh run_bench.sh
cd ..

echo "run rest benchmarks"
cd rest_client
sh run_bench.sh
cd ..

echo "run serialization benchmarks"
cd serialization
sh run_bench.sh
cd ..

python3 rest_vs_grpc_perf.py $1 $2
python3 serialization_perf.py $1 $2