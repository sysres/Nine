package sys

import (
	"runtime"

	"github.com/madlambda/nonix/sys/graphics"
)

type (
	Nonix struct {
		ScreenWidth  float64
		ScreenHeight float64
		Ctx2d        graphics.Ctx2d
	}
)

var (
	nonix *Nonix
)

// GoVersion returns the Go's runtime version.
func GoVersion() string {
	return runtime.Version()
}

// Bootstrap Nonix kernel
func Bootstrap() error {
	nonix := &Nonix{}
	err := bootstrap(nonix) // system-dependent
	if err != nil {
		return err
	}

	nonix.Ctx2d.ClearRect(0, 0, nonix.ScreenWidth, nonix.ScreenHeight)
	nonix.Ctx2d.SetFont("72px serif")
	nonix.Ctx2d.FillText("Welcome to Nonix Operating System", 10, 50)
	return nil
}
