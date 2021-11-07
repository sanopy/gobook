package writer

import "io"

type WriterWrapper struct {
	w   io.Writer
	cnt int64
}

func (c *WriterWrapper) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.cnt += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wrapper WriterWrapper
	wrapper.w = w
	return &wrapper, &wrapper.cnt
}
