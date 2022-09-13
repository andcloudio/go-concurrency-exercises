package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}

	}()

	// TODO: if there is no value on channel, do not block.
	for i := 0; i < 3; i++ {

		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			println("no message recieved")
		}

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)

	}
}
