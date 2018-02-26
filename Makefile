build:
	go build -o yd

build_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o yd_linux_amd64

build_linux_386:
	GOOS=linux GOARCH=386 go build -o yd_linux_386

build_linux_arm:
	GOOS=linux GOARCH=arm go build -o yd_linux_arm

build_windows_amd64:
	GOOS=windows GOARCH=amd64 go build -o yd_windows_amd64.exe

build_darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o yd_macOS

build_all:build_linux_amd64 build_linux_386 build_linux_arm build_windows_amd64 build_darwin_amd64
