package main

func main() {
	go func(a, b int) {
		c := a + b
	}(1, 2)
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
