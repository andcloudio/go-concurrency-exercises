package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer

	fmt.Fprintf(os.Stdout, "hello ")
	fmt.Fprintf(&buf, "world")
}
