package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	// get the value computed from goroutine
	c := <-ch
	fmt.Printf("computed value %v\n", c)
}
