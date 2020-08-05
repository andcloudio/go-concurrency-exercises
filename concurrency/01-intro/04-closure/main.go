package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Println(i)
		}()
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}
