package main

import "fmt"

func main() {
	//TODO: create pipeline with stages.
	// generator - generate values into channel [1, 2, 3, 4]
	// adder - add 2 to each element
	// multiplier - multiple by 3 each element
	// main - print output.

	done := make(chan interface{})
	defer close(done)

	intCh := generator(done, 1, 2, 3, 4)
	pipeline := multiplier(done, adder(done, intCh, 2), 3)

	for v := range pipeline {
		fmt.Println(v)
	}

}
