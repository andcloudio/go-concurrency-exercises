package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func(a, b int) {
		defer wg.Done()
		c := a + b
	}()

	wg.Wait()
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
