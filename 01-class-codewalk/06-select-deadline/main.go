package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "message"
	}()

	select {
	case m := <-ch:
		fmt.Println("received message", m)
	default:
		fmt.Println("no message received")
	}
}
