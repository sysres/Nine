package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Printf("hello")
	kchan := js.Global().Get("kchan")
	kchan.Call("postMessage", `{
		"type": "write",
		"obj": {
			"fd": 1,
			"data": "hello from wm",
			"count": 13
		}
	}`)
}
