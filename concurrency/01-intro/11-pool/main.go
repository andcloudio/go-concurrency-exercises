package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.

func log(w io.Writer, val string) {
	var b bytes.Buffer

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
