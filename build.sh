#!/usr/bin/env bash
export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin

protoc --go_out=plugins=grpc:. grpc/protos/*.proto
go build -o ./bin/dapi .