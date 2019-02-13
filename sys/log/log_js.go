package log

import "syscall/js"

func printf(str string) {
	js.Global().Get("console").Call("log", str)
}
