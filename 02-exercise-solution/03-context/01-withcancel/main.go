package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they consume 5 integers
	// so that internal goroutine
	// started by gen is not leaked.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	ch := gen(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
		}
	}
}
