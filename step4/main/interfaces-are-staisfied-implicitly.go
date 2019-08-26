package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Reader(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	//os.Stdout 实现类Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello,writer\n")

}
