package main

import (
	"io"
	"log"
)

func main() {
	// TODO: connect to server on localhost port 8000

}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
