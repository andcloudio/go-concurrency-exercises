package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			defer close(dst)

			numGenerator := func() error {
				deadline, ok := ctx.Deadline()
				if ok {
					if deadline.Sub(time.Now().Add(10*time.Millisecond)) <= 0 {
						return context.DeadlineExceeded
					}
				}
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case dst <- n:
						n++
					}
				}
			}

			err := numGenerator()
			if err != nil {
				fmt.Println(err)
			}
		}()
		return dst
	}

	deadline := time.Now().Add(5 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
	}
}
