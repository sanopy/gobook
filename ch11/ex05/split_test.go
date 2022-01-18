package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{
			s:    "",
			sep:  ",",
			want: 1,
		},
		{
			s:    "Hello, World!",
			sep:  ",",
			want: 2,
		},
		{
			s:    "Hello, World!",
			sep:  ".",
			want: 1,
		},
		{
			s:    "This is test message.",
			sep:  " ",
			want: 4,
		},
	}
	for _, tt := range tests {
		words := strings.Split(tt.s, tt.sep)
		if got := len(words); got != tt.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", tt.s, tt.sep, got, tt.want)
		}
	}
}
