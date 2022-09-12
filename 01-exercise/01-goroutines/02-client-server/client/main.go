package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// TODO: connect to server on localhost port 8000
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("error %w", err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
