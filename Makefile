.PHONY: build export run run-electron clean

SYSHOST=electron

ifdef SystemRoot
	EXPORT=set GOOS=js && set GOARCH=wasm
else
	EXPORT=export GOOS=js && export GOARCH=wasm
endif

all: clean build run

build: export
	go build -o view/$(SYSHOST)/dist/nine.wasm

export:
	$(shell $(EXPORT))

run: run-electron
run-electron:
	cd view/electron && make run

clean:
	rm view/$(SYSHOST)/dist/nine.wasm

