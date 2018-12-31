all: build deps

run: ./dist/wasm_exec.js
	npm start

dist/wasm_exec.js:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" $@

deps:
	sudo apt-get update && sudo apt-get install yarn
	npm install --save-dev electron
	yarn add electron-builder --dev

