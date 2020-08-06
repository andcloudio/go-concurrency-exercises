package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/dgunjetti/awesomeProject/concurrency/02-patterns/01-confinement/bank"
)

func main() {
	const GOMAXPROCS = 2
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup

	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			bank.Deposit(200)
		}()
		go func() {
			defer wg.Done()
			bank.Withdrawal(200)
		}()
	}

	wg.Wait()
	fmt.Println("=", bank.Balance())
}
