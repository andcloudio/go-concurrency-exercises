package main

import (
	"fmt"
	"sync"
)

func incr(wg *sync.WaitGroup) {
	var i int
	wg.Add(1)
	go func() {
		defer wg.Done()
		i++
		fmt.Println(i)
	}()
}

func main() {
	var wg sync.WaitGroup

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}
