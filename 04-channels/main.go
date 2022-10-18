package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type workFn[Result any] func(context.Context) Result

func main() {
	duration := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	dwf := func(ctx context.Context) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return "work complete"
	}
	result := doWork(ctx, dwf)
	select {
	case v := <-result:
		fmt.Println("main:", v)
	case <-ctx.Done():
		fmt.Println("main: timeout")
	}
}

func doWork[Result any](ctx context.Context, work workFn[Result]) chan Result {
	ch := make(chan Result, 1)
	go func() {
		ch <- work(ctx)
		fmt.Println("doWork : work complete")
	}()
	return ch
}
