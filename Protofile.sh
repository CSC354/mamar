#!/usr/bin/env sh

result=${PWD##*/}

protoc --go_out=. --go_opt=paths=import    \
    --go-grpc_out=. --go-grpc_opt=paths=import \
    protobufs/$result/*
