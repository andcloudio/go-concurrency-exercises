package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		for len(sharedRsc) == 0 {
			time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sharedRsc["rsc1"])
	}()

	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"

	wg.Wait()
}
