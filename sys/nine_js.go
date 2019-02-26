// +build js,wasm

package sys

import (
	"fmt"
	"syscall/js"
)

var beforeUnloadChan = make(chan struct{})

func bootstrap(nine *Nine) error {
	doc := js.Global().Get("document")
	nine.ScreenWidth = doc.Get("body").Get("clientWidth").Float()
	nine.ScreenHeight = doc.Get("body").Get("clientHeight").Float()
	return nil
}

func runtimeInfo() string {
	nodeVersion := js.Global().Get("process").Get("version")
	return fmt.Sprintf("NodeJS %s", nodeVersion.String())
}

func beforeUnload(this js.Value, args []js.Value) interface{} {
	beforeUnloadChan <- struct{}{}
	return nil
}

func wait() {
	beforeUnloadCb := js.FuncOf(beforeUnload)
	defer beforeUnloadCb.Release()

	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)
	<-beforeUnloadChan
}

func startService(name string) {
	js.Global().Call("loadService", name)
}
