GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# mac 系统更新path可能不全
export PATH := $(GOBIN):$(PATH)

macos-install:
	go install fyne.io/fyne/v2/cmd/fyne@latest #安装fyne命令行工具
	brew update -v
	brew install mingw-w64  #安装windows系统依赖

build-ico:
	fyne bundle yy.ico >> ./yyui/theme/icobundled.go #将静态资源编译为go文件

build-font:
	fyne bundle simkai.ttf >> ./yyui/theme/bundled.go #将静态资源编译为go文件

build-windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="x86_64-w64-mingw32-gcc" fyne package -os windows -icon yy.ico #打包windows二进制

build-macos:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 fyne package -os darwin -icon yy.ico #打包macos

build-all: macos-install build-windows build-macos

build-api:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./pinduoduo/server/  ./pinduoduo/server/main.go