package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// TODO: goroutine needs to cancelled after deadline 2 seconds from current time.

	var wg sync.WaitGroup
	gen := func() <-chan int {
		dst := make(chan int)
		n := 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(dst)
			for {
				select {
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	for n := range gen() {
		fmt.Println(n)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait()
}
