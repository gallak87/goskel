#!/usr/bin/env bash
## TODO: create make file to call this script

web=$(pwd)/../web
api=$web/src/api

# Server-side TS generation
#pbjs=$web/node_modules/.bin/pbjs
#pbts=$web/node_modules/.bin/pbts
#flags="--no-create --no-encode --no-decode --no-verify --no-convert --no-delimited --no-beautify"

# Clean previously generated files
rm -rf $api/*

# Generate server
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
  --grpc-web_out=import_style=typescript,mode=grpcweb:$api \
  ./userpb/*.proto

# Server-side TS generation
# $pbjs -p . $flags --no-comments -t static-module ./userpb/user.proto -o $api/user.service.js
# $pbjs -p . $flags -t static-module ./userpb/user.proto | \
#   $pbts --no-comments -o $api/user.service.d.ts - 

