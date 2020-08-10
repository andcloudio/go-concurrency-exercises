package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	//TODO: implement concurrent read.
	// allow multiple reads, writes holds the lock exclusively.

	wg.Wait()
	fmt.Println(balance)
}
