package main

import (
	"fmt"
	"time"

	"github.com/andcloudio/go-concurrency-exercises/01-intro/01-goroutines/04-add/counting"
)

func main() {
	numbers := counting.GenerateNumbers(1e7)

	t := time.Now()
	_ = counting.Add(numbers)
	fmt.Printf("Sequential Add took: %s\n", time.Since(t))

	t = time.Now()
	_ = counting.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add took: %s\n", time.Since(t))
}
