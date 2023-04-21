package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	// TODO: get the value computed from goroutine
	r := <-ch
	fmt.Printf("computed value %v\n", r)
}
