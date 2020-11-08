package main

import (
	"io"
	"os"
)

func main() {
	printer(os.Stdout, "hello")
}

func printer(w io.Writer, str string) {
	w.Write([]byte(str))
	w.Close()
}
