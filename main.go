package main

import (
	"fmt"
	"os"
	"time"

	"github.com/madlambda/Nine/sys"
)

func main() {
	err := sys.Bootstrap()
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}

	sys.Printf(5, 20, "Welcome to Nine Operating System!")
	sys.Printf(5, 50, "Go version: %s", sys.GoVersion())

	for {
		fmt.Printf(".")
		time.Sleep(time.Second * 2)
	}
}
