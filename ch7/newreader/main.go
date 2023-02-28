package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	if len(r.s) == 0 {
		return 0, io.EOF
	}
	n = copy(p, r.s)
	r.s = r.s[n:]
	return n, nil
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}

func main() {
	s := "https://golang.org"
	reader := NewReader(s)
	fmt.Printf("len(s) = %d\n", len(s))
	p := make([]byte, len(s))
	n, _ := reader.Read(p)
	fmt.Printf("Read %d bytes\n", n)
	n, _ = reader.Read(p)
	fmt.Printf("Read %d bytes\n", n)
}
