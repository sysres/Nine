// +build js,wasm

package sys

import (
	"syscall/js"
)

func bootstrap(nine *Nine) error {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "screen")

	nine.ScreenWidth = doc.Get("body").Get("clientWidth").Float()
	nine.ScreenHeight = doc.Get("body").Get("clientHeight").Float()
	nine.Ctx2d = GetJSCtx2d(canvasEl)
	return nil
}
