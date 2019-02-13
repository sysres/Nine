// +build js,wasm
package proc

import (
	"encoding/json"
	"fmt"

	"syscall/js"

	"github.com/madlambda/Nine/sys/log"
)

type (
	JSProc js.Value

	ProcMessage struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	jsFunc func(this js.Value, args []js.Value) interface{}
)

const workerScript = "proc.js"

func onError(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		panic("CRITICAL: onerror, no argument")
		return nil
	}

	err := args[0].Get("message").String()
	log.Printf(fmt.Sprintf("sched err: %s", err))
	return nil
}

func onProcMessage(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		panic("CRITICAL: onprocmessage, no argument")
		return nil
	}

	log.Printf("onprocmessage: %s", args[0].Get("data").String())
	return nil
}

func onLoadMessage(path string, worker js.Value) jsFunc {
	return func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			panic("CRITICAL: onmessage, no argument")
			return nil
		}

		data := args[0].Get("data").String()
		var msg ProcMessage
		err := json.Unmarshal([]byte(data), &msg)
		if err != nil {
			log.Printf("CRITICAL ERROR: unmarshal: %s", err)
			log.Printf("data: %s", data)
			return nil
		}

		if msg.Type != "status" {
			log.Printf("CRITICAL ERROR: unknown msg: %s", data)
			return nil
		}

		if msg.Value != "running" {
			log.Printf("ERROR: sched failed to execute proc")
			return nil
		}

		worker.Set("onmessage", js.FuncOf(onProcMessage))
		worker.Call("postMessage", fmt.Sprintf(`{
            "type": "load",
            "path": "%s"
		}`, path))

		return nil
	}
}

// TODO(i4k): make it sync, only return after process load successfully
func exec(path string) (JSProc, error) {
	worker := js.Global().Get("Worker").New(workerScript)
	worker.Set("onerror", js.FuncOf(onError))
	worker.Set("onmessage", js.FuncOf(onLoadMessage(path, worker)))
	return JSProc(worker), nil
}

func (p JSProc) Terminate() {
	js.Value(p).Call("close")
}
