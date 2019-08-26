package main

import "code.google.com/p/go-tour/reader"

type MyReader struct {
}

func main() {
	reader.Validate(MyReader{})
}
