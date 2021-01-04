package main

import (
  "context"
  "math/rand"
  "fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
    vals := make(chan int)
    
    go func(){
      defer close(vals)

      for {
        select {
          case vals <- rand.Int():
          case <-ctx.Done():
            return
        }
      }
    }()

    return vals
	}

	// Create a context that is cancellable.
  ctx, cancel := context.WithCancel(context.Background())

  receiver := generator(ctx)
  numList := []int{}

  for v := range receiver {
    numList = append(numList, v)

    if len(numList) == 5 {
      cancel()
    }
  }

  fmt.Printf("%#v\n", numList)
}
