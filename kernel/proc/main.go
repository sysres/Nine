package main

import (
	"fmt"
	"syscall/js"
)

func Hello(args []js.Value) { fmt.Printf("Hello from Proc") }

func main() {
	c := make(chan struct{})
	js.Global().Set("Hello", js.NewCallback(Hello))
	<-c
}
