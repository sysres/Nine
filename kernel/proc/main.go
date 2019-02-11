package main

import (
	"fmt"
	"syscall/js"

	"github.com/madlambda/Nine/sys"
)

func Hello(this js.Value, args []js.Value) interface{} {
	fmt.Printf("Hello from Proc")
	return nil
}

func main() {

	js.Global().Set("Hello", js.FuncOf(Hello))

	sys.Wait()
}
