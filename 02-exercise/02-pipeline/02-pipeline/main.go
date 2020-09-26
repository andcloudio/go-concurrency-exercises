// generator() -> square() -> print

package main

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
}

func main() {
	in := generator(2, 3)

	// TODO: fan out square stage to run two instances.

	// TODO: fan in the results of square stages.

}
