package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

//TODO: create pool of bytes.Buffers which can be reused.

var bufPool = sync.Pool{
	New: func() interface{} {
		println(" allocating new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
