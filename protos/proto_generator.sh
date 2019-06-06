#!/bin/bash

# go get google.golang.org/grpc
files=`ls *.proto`


protoc -I. -I/usr/local/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis $files --go_out=plugins=grpc:.

find ./ -name \*.pb.go -exec sed -i '' 's/import _ \"google\/api\"//g' {} \;
find ./ -name \*.pb.go -exec sed -i '' 's/protos\./grpc\/protos\./g' {} \;

protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true,stderrthreshold=0:. \
 $files

protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --swagger_out=logtostderr=true:. \
 $files






