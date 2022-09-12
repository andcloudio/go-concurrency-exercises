package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call

	go fun("go routine")

	// goroutine with anonymous function

	go func() {
		fun("anonymous function")
	}()

	// goroutine with function value call

	fv := fun

	go fv(" go routine 3")

	// wait for goroutines to end

	fmt.Println("done..")
	fmt.Println(" waiting for go routine")

	time.Sleep(1 * time.Second)

}
