GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# mac 系统更新path可能不全
export PATH := $(GOBIN):$(PATH)

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o  -o bin/yy ./
build-macos:
	go build -ldflags="-w -s" -o  -o bin/yy ./
build-windows:
    GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o bin/yy.exe ./