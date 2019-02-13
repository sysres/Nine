package main

import (
	"fmt"
	"os"

	"github.com/madlambda/Nine/sys"
)

func main() {
	err := sys.Bootstrap()
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}

	sys.Printf(5, 20, "Welcome to Nine OS!")
	sys.Printf(5, 40, sys.RuntimeInfo())
	sys.Wait()
}
