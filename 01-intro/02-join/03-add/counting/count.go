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
	runtime.GOMAXPROCS(runtime.NumCPU())

	numOfCores := runtime.NumCPU()
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

			// Add part sum to cummulative sum
			atomic.AddInt64(&sum, partSum)
		}(part)
	}

	wg.Wait()
	return sum
}
