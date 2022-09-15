package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.
	deadline := time.Now().Add(49 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)

			if deadline, ok := ctx.Deadline(); ok {
				if deadline.Sub(time.Now().Add(time.Millisecond*50)) < 0 {
					fmt.Print(" not sufficient time\n")
					return
				}
			}

			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			select {
			// Report result.
			case ch <- data{"123"}:
			case <-ctx.Done():
				return
			}

		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)

	// select {
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("overslept")
	// case <-ctx.Done():
	// 	fmt.Println(ctx.Err())
	// }

	if d, ok := <-ch; ok {
		fmt.Printf("work complete: %s\n", d)
	}

}
