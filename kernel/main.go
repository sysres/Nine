package main

import (
	"fmt"
	"os"

	"syscall/js"

	"github.com/madlambda/Nine/sys"
	"github.com/madlambda/Nine/sys/log"
	"github.com/madlambda/Nine/sys/proc"
)

var counter int

func onProcMessage(this js.Value, args []js.Value) interface{} {
	counter++
	log.Printf(args[0].Get("data").String())
	log.Printf("%d", counter)
	return nil
}

func main() {
	err := sys.Bootstrap()
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}

	sys.Printf(5, 20, "Welcome to Nine OS!")
	sys.Printf(5, 40, sys.RuntimeInfo())

	kchanProc := js.Global().Get("kchanProc")
	kchanProc.Set("onmessage", js.FuncOf(onProcMessage))

	sched := proc.NewSched()
	err = sched.Exec("wm/wm.wasm")
	if err != nil {
		sys.Printf(5, 120, "FATAL: %s", err)
		return
	}
	sys.Wait()
}
