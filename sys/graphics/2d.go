package graphics

type (
	Ctx2d interface {
		ClearRect(x1, y1, x2, y2 float64)
		SetFont(font string)
		FillText(str string, x, y float64)
	}
)
