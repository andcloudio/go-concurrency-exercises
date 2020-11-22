package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = time.Second
	fmt.Println(w)
}
