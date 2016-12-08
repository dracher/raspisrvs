SHELL = /bin/sh

stopraspsrv:
	ssh -p 2222 pi@dracher.mynetgear.com "supervisorctl stop raspsrv"

startraspsrv:
	ssh -p 2222 pi@dracher.mynetgear.com "supervisorctl start raspsrv"

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

local: dev static
	go run main.go bindata.go

localp: prod
	go run main.go bindata.go


cp2remote:
	scp -P 2222 build/raspisrv-arm pi@dracher.mynetgear.com:/home/pi/Projects/raspisrv

mergeui:
	cp /home/dracher/WebstormProjects/raspsrvui/dist/index.html templates/ && cp -r /home/dracher/WebstormProjects/raspsrvui/dist/static/* assets/

.PHONY: clean
clean:
	rm -rf assets/* && rm bindata.go build/* 

deploy: clean mergeui build-arm  stopraspsrv cp2remote startraspsrv
	echo "Deploy finished!"