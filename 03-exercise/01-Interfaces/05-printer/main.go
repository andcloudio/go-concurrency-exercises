package main

import (
	"os"
)

func main() {
	printer(os.Stdout, "hello")
}

func printer(f *os.File, str string) {
	f.Write([]byte(str))
	f.Close()
}

//	os.Stdout.Write([]byte("hello"))
//	os.Stdout.Close()
