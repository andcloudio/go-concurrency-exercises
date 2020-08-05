package main

import (
	"fmt"
	"sync"
)

func main() {
	// TODO: pass request scoped variables to handleResponse goroutine
	ProcessRequest("jane", "abc123")
}

func ProcessRequest(userID, authToken string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go handleResponse()
	wg.Wait()
}

func handleResponse() {
	fmt.Println("userID:")
	fmt.Println("authToken:")
}
