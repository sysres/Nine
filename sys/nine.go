package sys

import (
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

var (
	nine *Nine
)

// GoVersion returns the Go's runtime version.
func GoVersion() string {
	return runtime.Version()
}

// Bootstrap nine kernel
func Bootstrap() error {
	nine := &Nine{}
	err := bootstrap(nine) // system-dependent
	if err != nil {
		return err
	}

	nine.Ctx2d.ClearRect(0, 0, nine.ScreenWidth, nine.ScreenHeight)
	nine.Ctx2d.SetFont("72px serif")
	nine.Ctx2d.FillText("Welcome to nine Operating System", 10, 50)
	return nil
}
