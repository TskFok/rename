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

conf-local:

conf-rollback:

build-file:
	go build -o run -ldflags "-w -s"  -trimpath main.go

build-file-cli:
	go build -o run -ldflags "-w -s"  -trimpath bin/cli/main.go

build-mac: mac conf-local build-file conf-rollback

build-linux: linux conf-local build-file conf-rollback

build-cli-linux: linux conf-local build-file-cli conf-rollback

run-mac: mac
	go run main.go

update:
	go mody tidy
