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

func TestPopCountWithClearingLSBMinimumX(t *testing.T) {
	x := uint64(0)
	actual := PopCountWithClearingLSB(x)
	expected := 0
	if actual != expected {
		t.Errorf(`PopCountWithClearingLSB(%v) = %v`, x, actual)
	}
}

func TestPopCountWithClearingLSBMaximumX(t *testing.T) {
	x := uint64(0xffffffffffffffff)
	actual := PopCountWithClearingLSB(x)
	expected := 64
	if actual != expected {
		t.Errorf(`PopCountWithClearingLSB(%v) = %v`, x, actual)
	}
}

func BenchmarkPopCountWithExpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithExpr(0x012345679abcdef)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(0x012345679abcdef)
	}
}

func BenchmarkPopCountWithBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithBitwise(0x012345679abcdef)
	}
}

func BenchmarkPopCountWithClearingLSB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithClearingLSB(0x012345679abcdef)
	}
}
