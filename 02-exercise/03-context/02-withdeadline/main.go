package main

import (
	"fmt"
	"time"
  "context"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

  deadline := time.Now().Add(10 * time.Millisecond)
  ctx, cancel := context.WithDeadline(context.Background(), deadline)
  defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)

      deadl, ok := ctx.Deadline()
      if ok {
        if deadl.Sub(time.Now().Add(50*time.Millisecond)) < 0{
          fmt.Println("not sufficient time to do the process...")
          return
        }
      }
			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.
      select {
			  case ch <- data{"123"}:
        case <- ctx.Done():
      }
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)
	d, ok := <-ch
  if ok {
    fmt.Printf("work complete: %s\n", d)
  }
}
