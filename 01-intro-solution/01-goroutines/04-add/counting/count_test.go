package counting

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		AddConcurrent(numbers)
	}
}
