package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	read := func() int {
		mu.RLock()
		defer mu.RUnlock()
		return balance
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	// implement concurrent read.
	// allow multiple reads, writes holds the lock exclusively.
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(read())
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}
