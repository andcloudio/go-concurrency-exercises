package main

import (
	"fmt"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int

	go func() {
		data++
	}()

	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
