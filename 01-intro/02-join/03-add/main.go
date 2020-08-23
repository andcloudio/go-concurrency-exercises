package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgunjetti/sequential/counting"
)

func generateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

func main() {

	numbers := generateNumbers(1e7)

	t := time.Now()
	_ = counting.Add(numbers)
	fmt.Printf("Sequential Add took: %s\n", time.Since(t))

	t = time.Now()
	_ = counting.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add took: %s\n", time.Since(t))

}
