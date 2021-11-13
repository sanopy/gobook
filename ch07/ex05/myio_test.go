package myio

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLimitReaderReadOnce(t *testing.T) {
	s := "This is test message."
	lim := 10
	r := strings.NewReader(s)
	lr := LimitReader(r, int64(lim))
	size := 32

	t.Run("bellow the limit", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != int(lim) {
			t.Errorf("expected read bytes: %v, actual: %v", lim, n)
		}
		if err != nil {
			t.Errorf("read failed: %v", err)
		}

		want := make([]byte, size)
		copy(want, []byte("This is te"))
		if !bytes.Equal(buf, want) {
			t.Errorf("actual: %v, want: %v", buf, want)
		}
	})

	t.Run("exceeded the limit", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != 0 {
			t.Errorf("expected read bytes: %v, actual: %v", 0, n)
		}

		if err != io.EOF {
			t.Errorf("expected io.EOF")
		}
	})
}

func TestLimitReaderReadMoreThanOnce(t *testing.T) {
	s := "This is test message."
	r := strings.NewReader(s)
	lr := LimitReader(r, int64(11))
	size := 5

	t.Run("first read", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != int(size) {
			t.Errorf("expected read bytes: %v, actual: %v", size, n)
		}
		if err != nil {
			t.Errorf("read failed: %v", err)
		}

		want := make([]byte, size)
		copy(want, []byte("This "))
		if !bytes.Equal(buf, want) {
			t.Errorf("actual: %v, want: %v", buf, want)
		}
	})

	t.Run("second read", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != int(size) {
			t.Errorf("expected read bytes: %v, actual: %v", size, n)
		}
		if err != nil {
			t.Errorf("read failed: %v", err)
		}

		want := make([]byte, size)
		copy(want, []byte("is te"))
		if !bytes.Equal(buf, want) {
			t.Errorf("actual: %v, want: %v", buf, want)
		}
	})

	t.Run("third read", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != 1 {
			t.Errorf("expected read bytes: %v, actual: %v", 1, n)
		}
		if err != nil {
			t.Errorf("read failed: %v", err)
		}

		want := make([]byte, size)
		copy(want, []byte("s"))
		if !bytes.Equal(buf, want) {
			t.Errorf("actual: %v, want: %v", buf, want)
		}
	})

	t.Run("fourth read", func(t *testing.T) {
		buf := make([]byte, size)
		n, err := lr.Read(buf)
		if n != 0 {
			t.Errorf("expected read bytes: %v, actual: %v", 0, n)
		}

		if err != io.EOF {
			t.Errorf("expected io.EOF")
		}
	})
}
