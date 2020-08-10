package main

import (
	"fmt"
	"net/http"
	"sync"
)

//TODO: child goroutine to notify the main routine with result
// of http.get along with error status.

func main() {
	var wg sync.WaitGroup

	checkStatus :=
		func(done <-chan interface{}, urls ...string) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, url := range urls {
					resp, err := http.Get(url)
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Printf("url: %s Response: %v\n", url, resp.Status)
					select {
					case <-done:
						return
					default:
					}
				}
			}()
		}
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	checkStatus(done, urls...)
	wg.Wait()
}
