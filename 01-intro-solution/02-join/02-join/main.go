package main

import (
	"fmt"
	"runtime"
	"sync"
)

// what will be output

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	done := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		done = true
	}()
	wg.Wait()
	for !done {
	}
	fmt.Println("finished")
}
