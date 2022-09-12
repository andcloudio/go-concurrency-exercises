package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var wg sync.WaitGroup
	var data int
	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()
	fmt.Printf("the value of data is %v\n", data)
	wg.Wait()
	fmt.Println("Done..")
}
