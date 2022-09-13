package main

import (
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			println(v, " ch1")
		case v := <-ch2:
			println(v, " ch2")
		}
	}

	// TODO: multiplex recv on channel - ch1, ch2

}
