#!/usr/bin/env bash

# macos arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o mint main.go
tar czvf "evm-inscriptions_-macos-arm64".tar.gz mint settings.yml 使用说明.txt
rm -f mint

sleep 3

# macos amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o mint main.go
tar czvf "evm-inscriptions-macos-amd64".tar.gz mint settings.yml 使用说明.txt
rm -f mint

sleep 3

# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mint.exe main.go
tar czvf "evm-inscriptions-windows".tar.gz mint.exe settings.yml 使用说明.txt
rm -f mint.exe

sleep 3

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o mint main.go
tar czvf "evm-inscriptions-linux-amd64".tar.gz mint settings.yml 使用说明.txt
rm -f mint