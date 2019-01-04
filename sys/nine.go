package sys

import (
	"fmt"
	"runtime"

	"github.com/madlambda/Nine/sys/graphics"
)

type (
	Nine struct {
		ScreenWidth  float64
		ScreenHeight float64
		Ctx2d        graphics.Ctx2d
	}
)

const (
	defBgStyle   = "#000000"
	defFontStyle = "#ffffff"
	defFont      = "18px Helvetica"
)

var (
	nine *Nine
)

// GoVersion returns the Go's runtime version.
func GoVersion() string {
	return runtime.Version()
}

// Bootstrap nine kernel
func Bootstrap() error {
	nine = &Nine{}
	err := bootstrap(nine) // system-dependent
	if err != nil {
		return err
	}

	nine.Ctx2d.ClearRect(0, 0, nine.ScreenWidth, nine.ScreenHeight)
	nine.Ctx2d.SetFillStyle(defBgStyle)
	nine.Ctx2d.FillRect(0, 0, nine.ScreenWidth, nine.ScreenHeight)
	return nil
}

// Printf is a low level primitive to print formatted text to
// the screen. You need to pass the x and y position because
// at this point there's no concept of console yet.
func Printf(x, y float64, format string, args ...interface{}) {
	if nine.Ctx2d == nil {
		panic("Nine not initialized")
	}

	str := fmt.Sprintf(format, args...)
	nine.Ctx2d.SetFont(defFont)
	nine.Ctx2d.SetFillStyle(defFontStyle)
	nine.Ctx2d.FillText(str, x, y)
}
