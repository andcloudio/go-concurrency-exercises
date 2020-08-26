package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "one"
	}()

	// if there is no value on channel, do not block.

	select {
	case m := <-ch:
		fmt.Println(m)
	default:
		fmt.Println("no message in channel")
	}

}
