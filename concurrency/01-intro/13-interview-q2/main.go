package main

import (
	"math/rand"
	"sync"
)

const N = 10

//TODO: Fix the issue below to avoid "concurrent map writes" error.

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	println(len(m))
}
