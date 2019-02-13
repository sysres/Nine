package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Printf("hello")
	kchan := js.Global().Get("kchan")
	for i := 0; i < 100; i++ {
		kchan.Call("postMessage", fmt.Sprintf(`{
			"type": "write",
			"id": %d,
			"obj": {
				"fd": 1,
				"data": "hello from wm",
				"count": 13
			}
		}`, i))
	}
}
