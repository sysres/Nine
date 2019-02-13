package graphics

type (
	Ctx2d interface {
		ClearRect(x1, y1, x2, y2 float64)
		SetFont(font string)
		SetFillStyle(style string)
		FillText(str string, x, y float64)
		FillRect(x, y, width, height float64)
	}
)

func NewCtx2d() Ctx2d {
	return newCtx2d()
}
