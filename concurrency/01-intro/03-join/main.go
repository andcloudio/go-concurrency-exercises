package main

import (
	"fmt"
)

func main() {
	//TODO: modify the program to print
	// goroutine output "hello" deterministically.

	go func() {
		fmt.Println("hello")
	}()

	fmt.Println("Done..")
}
