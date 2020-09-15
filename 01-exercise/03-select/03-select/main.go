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
	for i := 0; i < 2; i++ {
		m := <-ch
		fmt.Println(m)

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)
	}
}
