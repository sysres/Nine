package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Welcome to nonix!\n")

	for {
		fmt.Printf(".")
		time.Sleep(time.Second * 2)
	}
}
