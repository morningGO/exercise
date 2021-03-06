#!/bin/bash

#set -x
export GOROOT=~/go
export PATH=$GOROOT/bin:$PATH

# go 1.2.1をダウンロードして展開
wget -q https://storage.googleapis.com/golang/go1.3.linux-amd64.tar.gz
tar -C ~/ -xzf go1.3.linux-amd64.tar.gz

make test

# すべてのファイルを実行
make run

# クロスコンパイルの準備
pushd ~/go/src
GOOS=windows GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
GOOS=darwin  GOARCH=amd64 ./make.bash --no-clean 2> /dev/null 1> /dev/null
popd


mkdir artifacts
pwd=`pwd`

# 各OS向けにbuild
for os in linux windows darwin; do
	GOOS=$os GOARCH=amd64 make build
	tar -C bin -cvzf $pwd/artifacts/${os}_amd64.tar.gz ${os}
done

