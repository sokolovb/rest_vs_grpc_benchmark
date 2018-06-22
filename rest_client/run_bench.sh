#!/usr/bin/env bash

go test -bench=.*_Proxy | awk '$1 ~ /Benchmark/ {print $1 " " $3}' | column -t > proxy.txt
go test -bench=.*_Pure | awk '$1 ~ /Benchmark/ {print $1 " " $3}' | column -t > pure.txt
