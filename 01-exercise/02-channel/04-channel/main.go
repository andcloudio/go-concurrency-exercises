package main

import (
	"fmt"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- int) {
	// send message on ch1

	ch1 <- 10

}

func relayMsg(ch1 <-chan int, ch2 chan<- int) {
	// recv message on ch1

	if value, ok := <-ch1; ok {
		ch2 <- value
	}

	// send it on ch2
}

func main() {
	// create ch1 and ch2

	ch1, ch2 := make(chan int), make(chan int)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// spine goroutine genMsg and relayMsg

	if value, ok := <-ch2; ok {
		fmt.Println(" message recieved ", value)
	}
	// recv message on ch2
}
