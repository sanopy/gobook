package rev

import (
	"testing"
)

func TestReverse(t *testing.T) {
	arg := [LENGTH]int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := arg
	reverse(&actual)
	expected := [LENGTH]int{8, 7, 6, 5, 4, 3, 2, 1}
	if actual != expected {
		t.Errorf("reverse(%v) = %v, want %v", arg, actual, expected)
	}
}
