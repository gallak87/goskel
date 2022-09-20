#!/usr/bin/env bash
## TODO: create make file to call this script

web=$(pwd)/../web
api=$web/src/api

# Clean previously generated files
rm -rf $api/*

# Generate new clients
# TODO: dynamic per pb eventually, hard coding userpb for now
protoc \
  -I./userpb \
  -I/usr/local/include/google \
  -I$HOME/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/google/api \
  --go_out=./userpb \
  --go_opt=paths=source_relative \
  --go-grpc_out=./userpb \
  --go-grpc_opt=paths=source_relative \
  --js_out=import_style=commonjs:$api \
  --plugin=protoc-gen-grpc-web=/usr/local/bin/protoc-gen-grpc-web \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcweb:$api \
  ./userpb/*.proto

## TODO: Look into PBTS https://github.com/protobufjs/protobuf.js/tree/master/cli/bin
