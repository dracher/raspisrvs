SHELL = /bin/sh

static:
	go-bindata assets/... templates/...

build: static
	go build -o build/raspisrv-x86 main.go bindata.go

build-arm: static
	
	env GOOS=linux GOARCH=arm go build -o build/raspisrv-arm main.go bindata.go

prod:
	sed -i 's/^dev.*/dev: no/g' conf.yml

dev:
	sed -i 's/^dev.*/dev: yes/g' conf.yml

local: dev
	go run main.go bindata.go

deploy: prod build-arm
	echo "done"

.PHONY: clean
clean:
	rm bindata.go build/*