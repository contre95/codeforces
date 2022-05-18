#!/bin/bash -xe

mkdir $1
pushd $1
go mod init $1
go mod tidy
touch input
popd
cat start.go > $1/main.go


