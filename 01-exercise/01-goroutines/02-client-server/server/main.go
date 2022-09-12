package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Errorf(" error in listening in 8000 %v", err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
