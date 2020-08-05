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

	// TODO: if there is no value on channel, do not block.

	m := <-ch
	fmt.Println(m)

}
