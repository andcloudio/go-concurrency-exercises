package main

import (
	"fmt"
	"runtime"
)

// what will be output
// How can I change this code so that the output is “finished”?

func main() {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")
}