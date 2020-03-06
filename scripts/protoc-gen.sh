#!/usr/bin/bash

# TODO: make sure gopath is set

function generate_for_for_version {
  PROTOFILES=api/proto/$1/*
  for f in $PROTOFILES; do
    mkdir -p api/swagger/$1
    mkdir -p pkg/api/$1
    protoc --proto_path=api/proto/$1 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:pkg/api/$1 $(basename $f)
    protoc --proto_path=api/proto/$1 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:pkg/api/$1 $(basename $f)
    protoc --proto_path=api/proto/$1 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:api/swagger/$1 $(basename $f)
  done
}

for d in $(find api/proto/* -type d); do
  generate_for_for_version $(basename $d)
done
