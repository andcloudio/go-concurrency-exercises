package main

import "fmt"

func main() {
	describe(42)
	describe("hello")
}

func describe(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("v is integer with value %d\n", v)
	case string:
		fmt.Printf("v is a string, whose length is %d\n", len(v))
	default:
		fmt.Println("we dont know what 'v' is!")
	}
}
