package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	generator := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)
			for {
				select {
				case <-ctx.Done():
					log.Println(ctx.Err())
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

	ch := generator(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
		}
	}
}
