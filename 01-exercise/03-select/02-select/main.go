package main

import (
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch

	select {
	case m := <-ch:
		println(m)
	case <-time.After(time.Second * 1):
		println(" timeout")
	}

}
