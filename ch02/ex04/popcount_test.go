package popcount

import "testing"

func TestPopCountWithExprMinimumX(t *testing.T) {
	x := uint64(0)
	actual := PopCountWithExpr(x)
	expected := 0
	if actual != expected {
		t.Errorf(`PopCountWithExpr(%v) = %v`, x, actual)
	}
}

func TestPopCountWithExprMaximumX(t *testing.T) {
	x := uint64(0xffffffffffffffff)
	actual := PopCountWithExpr(x)
	expected := 64
	if actual != expected {
		t.Errorf(`PopCountWithExpr(%v) = %v`, x, actual)
	}
}

func TestPopCountWithLoopMinimumX(t *testing.T) {
	x := uint64(0)
	actual := PopCountWithLoop(x)
	expected := 0
	if actual != expected {
		t.Errorf(`PopCountWithLoop(%v) = %v`, x, actual)
	}
}

func TestPopCountWithLoopMaximumX(t *testing.T) {
	x := uint64(0xffffffffffffffff)
	actual := PopCountWithLoop(x)
	expected := 64
	if actual != expected {
		t.Errorf(`PopCountWithLoop(%v) = %v`, x, actual)
	}
}

func TestPopCountWithBitwiseMinimumX(t *testing.T) {
	x := uint64(0)
	actual := PopCountWithBitwise(x)
	expected := 0
	if actual != expected {
		t.Errorf(`PopCountWithBitwise(%v) = %v`, x, actual)
	}
}

func TestPopCountWithBitwiseMaximumX(t *testing.T) {
	x := uint64(0xffffffffffffffff)
	actual := PopCountWithBitwise(x)
	expected := 64
	if actual != expected {
		t.Errorf(`PopCountWithBitwise(%v) = %v`, x, actual)
	}
}

var input uint64 = 0x012345679abcdef
var output int

func BenchmarkPopCountWithExpr(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithExpr(input)
	}
	output = s
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithLoop(input)
	}
	output = s
}

func BenchmarkPopCountWithBitwise(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithBitwise(input)
	}
	output = s
}
