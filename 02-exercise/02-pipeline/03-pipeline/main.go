// generator() -> square() ->
//														-> merge -> print
//             -> square() ->
package main

import (
	"fmt"
	"sync"
  "runtime"
  "time"
)

func generator(cancel <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
    defer close(out)
		for _, n := range nums {
      select {
        case out <- n:
        case <- cancel:
          return
      }
		}
	}()
	return out
}

func square(cancel <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
    defer close(out)
		for n := range in {
      select {
        case out <- n * n:
        case <- cancel:
          return
      }
		}
	}()
	return out
}

func merge(cancel <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
    defer wg.Done()
		for n := range c {
			select {
        case out <- n:
        case <- cancel:
          return
      }
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
  cancel := make(chan struct{})
	in := generator(cancel, 2, 3)

	c1 := square(cancel, in)
	c2 := square(cancel, in)

	out := merge(cancel, c1, c2)

	// TODO: cancel goroutines after receiving one value.

	fmt.Println(<-out)
  close(cancel)

  time.Sleep(10*time.Millisecond)
  fmt.Printf("Running goroutines: %v\n", runtime.NumGoroutine())
}
