package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	// TODO: range over channel to recv values

}
