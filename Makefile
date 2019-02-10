.PHONY: build export run run-electron clean

SYSHOST=electron

ifdef SystemRoot
EXPORT=set "GOOS=js"&&set "GOARCH=wasm"
else
EXPORT=export GOOS=js && export GOARCH=wasm
endif

all: clean build run

build:
	$(shell $(EXPORT) && cd kernel/ && go build -o ../view/$(SYSHOST)/dist/kern.wasm)
	$(shell $(EXPORT) && cd kernel/proc && go build -o ../../view/$(SYSHOST)/dist/proc.wasm)

run: run-electron
run-electron:
	cd view/electron && make run

clean:
	rm -f view/$(SYSHOST)/dist/kern.wasm
	rm -f view/$(SYSHOST)/dist/proc.wasm

