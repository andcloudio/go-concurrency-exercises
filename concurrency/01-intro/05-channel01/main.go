package main

func main() {
	// TODO: send message from goroutine to main goroutine.
	go func() {
		msg := "hi from goroutine"
	}()
}
