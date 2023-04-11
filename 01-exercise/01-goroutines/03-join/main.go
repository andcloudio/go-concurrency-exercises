package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		data++
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
