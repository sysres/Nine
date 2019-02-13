package sys

import (
	"fmt"
	"runtime"
	"strings"

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
	defBgStyle    = "#000000"
	defFontStyle  = "#ffffff"
	defFontSize   = 18 // in pixels
	defFontFamily = "Helvetica"

	NineVersion = "0.1"
)

var (
	nine    *Nine
	defFont string
)

func init() {
	defFont = fmt.Sprintf("%dpx %s", defFontSize, defFontFamily)
}

// Version of Nine
func Version() string {
	return fmt.Sprintf("Nine v%s", NineVersion)
}

// RuntimeInfo shows information about the various runtimes of Nine
func RuntimeInfo() string {
	return fmt.Sprintf(`%s
Go version: %s
Runtime info: %s`, Version(), GoVersion(), runtimeInfo())
}

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

	nine.Ctx2d = graphics.NewCtx2d()

	ClearScreen()
	return nil
}

func ClearScreen() {
	nine.Ctx2d.ClearRect(0, 0, nine.ScreenWidth, nine.ScreenHeight)
	nine.Ctx2d.SetFillStyle(defBgStyle)
	nine.Ctx2d.FillRect(0, 0, nine.ScreenWidth, nine.ScreenHeight)
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

	txts := strings.Split(str, "\n")
	for i, txt := range txts {
		nine.Ctx2d.FillText(txt, x, y+float64(i*(defFontSize+1)))
	}
}

func StartService(name string) {
	startService(name)
}

func Wait() {
	wait()
}
