package main
import (
  "math/rand"
  "sync"
  "fmt"
  "time"
)

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int)  <-chan int {
  out := make(chan int, len(nums))

  go func(){
    for _, n := range nums {
      fmt.Printf("1)\t saving %v\n", n)
      out <- n
    }

    close(out)
  }()

  return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(chIn <-chan int) <-chan int {
  out := make(chan int)

  go func(){
    for n := range chIn {
      fmt.Printf("2)\t %v*%v = %v\n", n, n, n*n)
      out <- n*n
    }

    close(out)
  }()

  return out
}

func main() {
  start := time.Now()
	// set up the pipeline
  vals := []int{}
  for i := 0; i < 100; i++ {
    vals = append(vals, rand.Intn(99)+1)
  }

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func(){
    for v := range square(generator(vals...)) {
      fmt.Printf("3)\t result: %v\n", v)
    }

    wg.Done()
  }()

  wg.Wait()
  fmt.Printf("elapsed time: %vms\n", time.Since(start).Milliseconds())
}
