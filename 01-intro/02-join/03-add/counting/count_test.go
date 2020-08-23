package counting

import (
	"math/rand"
	"testing"
)

func generateNumbers(max int) []int {
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

func BenchmarkAdd(b *testing.B) {
	numbers := generateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	numbers := generateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		AddConcurrent(numbers)
	}
}
