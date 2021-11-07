package writer

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, n := CountingWriter(os.Stdout)
	t.Run(`write "Hello"`, func(t *testing.T) {
		w.Write([]byte("Hello"))
		want := int64(5)
		if *n != want {
			t.Errorf("actual: %v, want: %v", *n, want)
		}
	})
	t.Run(`write "World"`, func(t *testing.T) {
		w.Write([]byte("World"))
		want := int64(10)
		if *n != want {
			t.Errorf("actual: %v, want: %v", *n, want)
		}
	})
}
