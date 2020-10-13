package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {

	// Utilize all cores on machine
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var sum int64
	max := len(numbers)

	sizeOfParts := max / numOfCores

	var wg sync.WaitGroup

	for i := 0; i < numOfCores; i++ {

		// Divide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]

		// Run computation for each part in seperate goroutine.
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()

			var partSum int64

			// Calculate sum for each part
			for _, n := range nums {
				partSum += int64(n)
			}

			// Add sum of each part to cummulative sum
			atomic.AddInt64(&sum, partSum)
		}(part)
	}

	wg.Wait()
	return sum
}
