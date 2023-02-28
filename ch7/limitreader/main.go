package main

import "io"

type LReader struct {
	r io.Reader
	n int64
}

func (r *LReader) Read(p []byte) (int, error) {
	if r.n == 0 {
		return 0, io.EOF
	}
	m, err := r.r.Read(p[:r.n-1])
	r.n -= int64(m)
	return m, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{r: r, n: n}
}
