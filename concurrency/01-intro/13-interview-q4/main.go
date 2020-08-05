package main

const NumberOfReaders = 10
const NumberOfWriters = 10

//TODO: complete the program to make writer and readers
// concurrency safe.

func writer() {
	// Write to data.
}

func reader() {
	// Read from data.
}

func main() {
	for i := 0; i < NumberOfReaders; i++ {
		go reader()
	}
	for i := 0; i < NumberOfWriters; i++ {
		go writer()
	}
}
