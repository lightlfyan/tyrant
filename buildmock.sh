#!/bin/bash

OLD=$GOPATH
CURR=`pwd`

export GOPATH=$CURR:$OLD
gofmt -w ./src/mock/*
#go build -o ./bin/mock -gcflags='-m' ./src/mock/main.go 2>&1
go build -o ./bin/mock ./src/mock/main.go
cp ./etc/conf.json ./bin/
export GOPATH=$OLD
