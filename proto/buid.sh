#!/bin/bash
DIR=$(cd $(dirname $0); pwd)
cd $DIR

protoc --go_out=. *.proto
cp pb.pb.go ../src/master/protocol/pb.go
mv pb.pb.go ../src/cell/protocol/pb.go
