package main

import (
	"fmt"
	"os"

	"syscall/js"

	"github.com/madlambda/Nine/sys"
)

var beforeUnloadChan = make(chan struct{})

func MouseMove(args []js.Value) {
	evt := args[0]
	x := evt.Get("x").Float()
	y := evt.Get("y").Float()

	sys.ClearScreen()
	sys.Printf(5, 100, "x=%f, y=%f\n", x, y)
}

func beforeUnload(event js.Value) {
	fmt.Printf("BeforeUNload fired\n")
	beforeUnloadChan <- struct{}{}
}

func main() {
	err := sys.Bootstrap()
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}

	sys.Printf(5, 20, "Kernel loaded\n")

	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	defer beforeUnloadCb.Release()

	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	cb := js.NewCallback(MouseMove)
	defer cb.Release()

	js.Global().Get("document").Get("body").Call("addEventListener", "mousemove", cb)

	<-beforeUnloadChan
}
