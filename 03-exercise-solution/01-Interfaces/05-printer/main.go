package main

import (
	"io"
	"os"
)

func main() {
	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	printer(os.Stdout)
}

func printer(w io.Writer) {
	w.Write([]byte("hello"))
	w.Close()
}
