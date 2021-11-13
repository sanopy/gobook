package myio

import (
	"io"
	"math"
)

type LimitedReader struct {
	r   io.Reader
	i   int64
	lim int64
}

func (lr *LimitedReader) Read(b []byte) (n int, err error) {
	if lr.i >= lr.lim {
		return 0, io.EOF
	}
	end := int(math.Min(float64(len(b)), float64(lr.lim-lr.i)))
	n, err = lr.r.Read(b[:end])
	lr.i += int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, 0, n}
}
