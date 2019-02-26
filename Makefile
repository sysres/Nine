.PHONY: build export run run-electron clean build-image local-build

ifdef SystemRoot
CWD=$(shell echo %cd%)
else
CWD=$(shell pwd)
endif

SYSHOST=electron
IMAGE=madlambda/nine-build-tools:0.1
RUN=docker run -it --rm -v $(CWD):/go/src/github.com/madlambda/Nine $(IMAGE)

all: clean build run
build: build-image
	$(RUN) make local-build

test: build-image
	$(RUN) make local-test

deps:
	$(RUN) go get -v ./...

shell:
	$(RUN) bash

local-build:
	cd kernel/ && go build -o ../view/$(SYSHOST)/dist/kern.wasm
	cd sys/cmd/wm && go build -o ../../../view/$(SYSHOST)/dist/wm/wm.wasm

local-test:
	go test -v ./...

run: run-electron
run-electron:
	cd view/electron && make run

build-image:
	docker build -t $(IMAGE) .

clean:
	rm -f view/$(SYSHOST)/dist/kern.wasm
	rm -f view/$(SYSHOST)/dist/wm/wm.wasm

