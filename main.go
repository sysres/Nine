package main

import (
	"fmt"
	"os"
	"time"

	"github.com/madlambda/Nine/sys"
)

func main() {

	fmt.Printf("Welcome to nonix!\n")
	fmt.Printf("Go version: %s\n", sys.GoVersion())

	err := sys.Bootstrap()
	if err != nil {
		fmt.Printf("FATAL: %s\n", err)
		os.Exit(1)
	}

	for {
		fmt.Printf(".")
		time.Sleep(time.Second * 2)
	}
}
