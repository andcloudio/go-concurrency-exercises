package main

import "fmt"

func main() {
	describe(42)
	describe("hello")
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
