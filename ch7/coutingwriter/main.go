package main

import "io"

type CWriter struct {
	inner io.Writer
	count int64
}

func (w *CWriter) Write(p []byte) (n int, err error) {
	n, err = w.inner.Write(p)
	w.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CWriter{inner: w, count: 0}
	return &cw, &cw.count
}
