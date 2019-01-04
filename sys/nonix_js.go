// +build js,wasm

package sys

import (
	"syscall/js"
)

func bootstrap(nonix *Nonix) error {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "screen")

	nonix.ScreenWidth = doc.Get("body").Get("clientWidth").Float()
	nonix.ScreenHeight = doc.Get("body").Get("clientHeight").Float()

	canvasEl.Set("width", nonix.ScreenWidth)
	canvasEl.Set("height", nonix.ScreenHeight)
	canvasEl.Set("style", "cursor: none")

	nonix.Ctx2d = GetJSCtx2d(canvasEl)
	return nil
}
