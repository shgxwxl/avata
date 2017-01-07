GO ?= go 
PWD := $(shell pwd)
GOPATH := $(GOPATH):$(PWD)/_3rd:$(PWD)/src:$(PWD)

all : build

build : init
	export GOPATH=$(GOPATH)
	$(GO) build  -o bin/avata_admin

gdb : init
	export GOPATH=$(GOPATH):$(PWD)/../../
check :
	echo check
	go test ./push/
init :
	echo init
clean:
	rm -rf bin/avata_admin