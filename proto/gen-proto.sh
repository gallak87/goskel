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
#   --plugin=protoc-gen-ts=$web/node_modules/.bin/protoc-gen-ts \
#   --ts_opt=unary_rpc_promise=true \
#   --ts_opt=json_names \
#   --ts_opt=explicit_override \
#   --ts_out=service=grpc-web:$api \
#  ./userpb/*.proto

## FOR TS-PROTO ONLY
#   --plugin=protoc-gen-ts_proto=$web/node_modules/.bin/protoc-gen-ts_proto \
#   --ts_proto_out=. \
#   --ts_proto_opt=esModuleInterop=true \
#   --ts_proto_opt=returnObservable=false \
#   ./userpb/*.proto
## TODO: bug in path, --ts_proto_out=$api results in gallak87/goskel/proto/../web/src/api/: No such file or directory
## -> just copy the output 
# mv *.[t,j]s $api/
## TOOD:    --ts_proto_opt=exportCommonSymbols=false,unknownFields=true,usePrototypeForDefaults=true \