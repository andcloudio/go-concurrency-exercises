package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 3)

	wg.Add(2)
	go G1(ch)
	go G2(ch)
	wg.Wait()
}

//G1 - goroutine
func G1(ch chan<- int) {
	defer wg.Done()
	for _, v := range []int{1, 2, 3, 4} {
		ch <- v
	}
	close(ch)
}

//G2 - goroutine
func G2(ch <-chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
