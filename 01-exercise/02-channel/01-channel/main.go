package main

import "fmt"

func main() {

	ch := make(chan int)

	go func(a, b int, ch chan int) {
		c := a + b
		ch <- c
	}(1, 2, ch)

	fmt.Println(<-ch)
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
