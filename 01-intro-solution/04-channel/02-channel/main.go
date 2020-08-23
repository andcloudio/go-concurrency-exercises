package main

func main() {
	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
		}
	}()

	// TODO: range over channel to recv values

}
