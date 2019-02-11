// +build js,wasm

package sys

import (
	"syscall/js"
)

var beforeUnloadChan = make(chan struct{})

func beforeUnload(this js.Value, args []js.Value) interface{} {
	beforeUnloadChan <- struct{}{}
	return nil
}

func bootstrap(nine *Nine) error {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "screen")

	nine.ScreenWidth = doc.Get("body").Get("clientWidth").Float()
	nine.ScreenHeight = doc.Get("body").Get("clientHeight").Float()
	nine.Ctx2d = GetJSCtx2d(canvasEl)
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
