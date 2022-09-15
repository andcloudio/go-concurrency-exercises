// generator() -> square() -> print

package main

import (
	"sync"
	"time"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(2 * time.Second)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel

	println("starting merge channel")

	out := make(chan int)
	sendToOut := func(c <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go sendToOut(c, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
		println(" closing as read is done")
	}()

	return out
}

func main() {
	in := generator(1, 2, 3, 4)

	// TODO: fan out square stage to run two instances.
	ch1 := square(in)
	ch2 := square(in)

	// TODO: fan in the results of square stages.
	merge := merge(ch1, ch2)

	for v := range merge {
		println(v)
	}
}
