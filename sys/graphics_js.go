package sys

import "syscall/js"

type (
	JSCtx2d struct {
		js.Value
	}
)

func GetJSCtx2d(canvasElement js.Value) *JSCtx2d {
	return &JSCtx2d{
		Value: canvasElement.Call("getContext", "2d"),
	}
}

func (ctx *JSCtx2d) ClearRect(x1, y1, x2, y2 float64) {
	ctx.Call("clearRect", x1, y1, x2, y2)
}

func (ctx *JSCtx2d) SetFont(font string) {
	ctx.Set("Font", font)
}

func (ctx *JSCtx2d) FillText(str string, x, y float64) {
	ctx.Call("fillText", str, x, y)
}
