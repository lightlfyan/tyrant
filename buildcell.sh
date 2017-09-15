#!/bin/bash

OLD=$GOPATH
CURR=`pwd`

export GOPATH=$CURR:$OLD
goimports -w ./src
gofmt -w ./src/cell/*
go build -o ./bin/cell ./src/cell/main.go
export GOPATH=$OLD
