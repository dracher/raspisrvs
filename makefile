SHELL = /bin/sh

static:
	go-bindata assets/... templates/...

build: static
	go build -o build/raspisrv-x86 main.go bindata.go

build-arm: static
	env GOOS=linux GOARCH=arm go build -o build/raspisrv-arm main.go bindata.go

prod: build
	./build/raspisrv-x86

.PHONY: clean
clean:
	rm bindata.go build/*

dev:
	env DEV=1 go run main.go bindata.go