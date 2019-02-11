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

	sys.Printf(5, 20, "Kernel loaded\n")
	sys.Wait()
}
