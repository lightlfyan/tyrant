#!/bin/bash
OLD=$GOPATH
CURR=`pwd`

export GOPATH=$CURR:$OLD
#goimports -w ./src/
gofmt -w ./src/master/*
#go build -race -o ./bin/master ./src/master/main.go
go build -o ./bin/master ./src/master/main.go
export GOPATH=$OLD
