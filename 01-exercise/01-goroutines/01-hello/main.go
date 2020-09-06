package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	f("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call

	// goroutine with anonymous function

	// goroutine with function value call

	// wait for goroutines to end

	fmt.Println("done..")
}
