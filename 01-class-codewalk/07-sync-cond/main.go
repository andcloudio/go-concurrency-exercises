package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]string)

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for len(sharedRsc) == 0 {
			mu.Unlock()
			time.Sleep(1 * time.Millisecond)
			mu.Lock()
		}
		fmt.Println(sharedRsc["rsc1"])
		mu.Unlock()
	}()

	mu.Lock()
	sharedRsc["rsc1"] = "foo"
	mu.Unlock()

	wg.Wait()
}
