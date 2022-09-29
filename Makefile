GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# mac 系统更新path可能不全
export PATH := $(GOBIN):$(PATH)

install:
	go install fyne.io/fyne/v2/cmd/fyne@latest #安装fyne命令行工具

build-ico:
	fyne bundle yy.ico >> ./yyui/theme/icobundled.go #将静态资源编译为 go 文件

build-windows:
	fyne package -os windows -icon yy.ico

build-macos:
	/Users/xiewenbin/go/src/bin/fyne package -os darwin -icon yy.ico
