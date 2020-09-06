package main

import (
	"context"
	"fmt"
)

func main() {

	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel() // cancel when done with consuming integers.
		}
	}
}
