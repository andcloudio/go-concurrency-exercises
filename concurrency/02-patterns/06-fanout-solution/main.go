package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
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

	fanIn := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		// Select from all the channels
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand := func() int { return rand.Intn(50000000) }

	randIntStream := generatorFn(done, rand)

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
