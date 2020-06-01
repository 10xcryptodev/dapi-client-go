#!/usr/bin/env bash
GO_LOCATION=$(which go)
GOROOT=$(echo ${GO_LOCATION%/bin/go})

protoc --go_out=plugins=grpc:. grpc/protos/*.proto
go build -o ./bin/dapi .