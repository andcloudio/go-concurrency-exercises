package main

import "fmt"

// TODO: Implement ping pong with Channel Direction

func ping(out chan<- string) {
	// send message on ch1
	out <- "ping"
}

func pong(in <-chan string, out chan<- string) {
	// recv message on ch1
	msg := <-in
	msg = msg + " pong"
	// send it on ch2
	out <- msg
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine ping and pong
	go ping(ch1)
	go pong(ch1, ch2)

	// recv message on ch2
	msg := <-ch2

	fmt.Println(msg)
}
