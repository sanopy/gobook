package mystrings

import (
	"io"
)

type Reader struct {
	s string
	i int
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func NewReader(s string) *Reader {
	return &Reader{s, 0}
}
