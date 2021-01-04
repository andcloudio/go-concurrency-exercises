// generator() -> square() -> print

package main

import (
  "fmt"
  "math/rand"
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
			out <- n * n
      time.Sleep(time.Duration(rand.Intn(500)+500)*time.Millisecond)
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
  out := make(chan int)
  wg := sync.WaitGroup{}

  for _, chInput := range cs {
    wg.Add(1)
    go func(ch <-chan int, out chan<- int){
      for val := range ch {
        out <- val
      }

      wg.Done()
    }(chInput, out)
  }

  go func(){
    wg.Wait()
    close(out)
  }()

  return out
}

func main() {
  content := []int{}
  for i := 0; i < 50; i++ {
    v := rand.Intn(100)+1
    fmt.Printf("Adding %v\n", v)
    content = append(content, v)
  }

  in := generator(content...)

	// TODO: fan out square stage to run two instances.
  ch1 := square(in)
  ch2 := square(in)
  ch3 := square(in)

	// TODO: fan in the results of square stages.
  for v := range merge(ch1, ch2, ch3) {
    fmt.Println(v)
  }
}
