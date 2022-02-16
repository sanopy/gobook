// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	w   io.WriteCloser
	cmd *exec.Cmd
	sync.Mutex
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("/bin/bzip2")
	cmd.Stdout = out

	in, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	w := &writer{w: in, cmd: cmd}
	return w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	return w.w.Write(data)
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()

	if err := w.w.Close(); err != nil {
		return err
	}
	return w.cmd.Wait()
}
