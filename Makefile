SYSHOST=electron

all: build run

build:
	GOOS=js GOARCH=wasm go build -o view/$(SYSHOST)/dist/nine.wasm

run: run-electron
run-electron:
	cd view/electron && make run

