package sys

import "syscall/js"

type (
	JSCtx2d struct {
		canvas js.Value
		js.Value
	}
)

func GetJSCtx2d(canvasEl js.Value) *JSCtx2d {
	return &JSCtx2d{
		canvas: canvasEl,
		Value:  canvasEl.Call("getContext", "2d"),
	}
}

func (ctx *JSCtx2d) ClearRect(x1, y1, x2, y2 float64) {
	ctx.Call("clearRect", x1, y1, x2, y2)
}

func (ctx *JSCtx2d) SetFont(font string) {
	ctx.Set("font", font)
}

func (ctx *JSCtx2d) FillText(str string, x, y float64) {
	ctx.Call("fillText", str, x, y)
}

func (ctx *JSCtx2d) FillRect(x, y, width, height float64) {
	ctx.Call("fillRect", x, y, width, height)
}

func (ctx *JSCtx2d) SetFillStyle(style string) {
	ctx.Set("fillStyle", style)
}
