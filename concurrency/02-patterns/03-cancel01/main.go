package main

import (
	"fmt"
	"sync"
)

// TODO: add the logic to cancel the goroutine

func main() {
	var wg sync.WaitGroup

	doWork := func(
		strings <-chan string,
	) {
		defer wg.Done()
		defer fmt.Println("doWork exited")
		for {
			select {
			case s := <-strings:
				fmt.Println(s)
			}
		}
	}

	wg.Add(1)
	go doWork(nil)

	wg.Wait()
	fmt.Println("done...")
}
