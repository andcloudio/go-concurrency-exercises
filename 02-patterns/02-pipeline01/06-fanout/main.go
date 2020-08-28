package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generatorFn := func(
		done <-chan interface{},
		fn func() int,
	) <-chan int {
		valueStream := make(chan int)
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	primeFinder :=
		func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
			primeStream := make(chan interface{})
			go func() {
				defer close(primeStream)
				for integer := range intStream {
					integer -= 1
					prime := true
					for divisor := integer - 1; divisor > 1; divisor-- {
						if integer%divisor == 0 {
							prime = false
							break
						}
					}

					if prime {
						select {
						case <-done:
							return
						case primeStream <- integer:
						}
					}
				}
			}()
			return primeStream
		}

	rand := func() int { return rand.Intn(50000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := generatorFn(done, rand)

	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
