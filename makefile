mac:
	go env -w GOARCH=amd64
	go env -w GOOS=darwin
	go env -w CGO_ENABLED=0
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod  tidy

linux:
	go env -w GOARCH=amd64
	go env -w GOOS=linux
	go env -w CGO_ENABLED=0
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod  tidy

build-file-cli:
	go build -o rename -ldflags "-w -s"  -trimpath bin/cli/main.go

build-cli-linux: linux build-file-cli

build-cli-mac: mac build-file-cli

run-mac: mac
	go run main.go

update:
	go mody tidy
