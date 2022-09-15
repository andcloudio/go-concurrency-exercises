package main

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range nums {
			ch <- v
		}
	}()

	return ch
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func main() {
	// set up the pipeline

	for v := range square(generator(1, 2, 3)) {
		println(v)
	}

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
