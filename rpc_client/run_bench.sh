#!/usr/bin/env bash

go test -bench=. | awk '$1 ~ /Benchmark/ {print $1 " " $3}' | column -t > rpc.txt
