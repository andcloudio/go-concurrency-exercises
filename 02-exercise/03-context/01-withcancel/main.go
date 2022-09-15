package main

import (
	"context"
	"sync"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.

	ctx, cancel := context.WithCancel(context.Background())

	generator := func(ctx context.Context) <-chan int {

		out := make(chan int)
		i := 1
		go func() {
			defer close(out)
			for {
				select {
				case out <- i:
					i++
				case <-ctx.Done():
					return
				}

			}
		}()

		return out
	}

	in := generator(ctx)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer func() {
			wg.Done()
		}()
		for v := range in {
			println(v)
			if v == 5 {
				cancel()
				return
			}
		}
	}(ctx)

	wg.Wait()

	// Create a context that is cancellable.

}
