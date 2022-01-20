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

var output int

func benchmarkPopCountWithExpr(b *testing.B, in uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithExpr(in)
	}
	output = s
}

func benchmarkPopCountWithLoop(b *testing.B, in uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithLoop(in)
	}
	output = s
}

func benchmarkPopCountWithBitwise(b *testing.B, in uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithBitwise(in)
	}
	output = s
}

func benchmarkPopCountWithClearingLSB(b *testing.B, in uint64) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountWithClearingLSB(in)
	}
	output = s
}

var expect16 uint64 = 0xFFFF
var expect32 uint64 = 0xFFFFFFFF
var expect64 uint64 = 0xFFFFFFFFFFFFFFFF

func BenchmarkPopCountWithExprExpect0(b *testing.B) {
	benchmarkPopCountWithExpr(b, 0)
}
func BenchmarkPopCountWithExprExpect16(b *testing.B) {
	benchmarkPopCountWithExpr(b, expect16)
}
func BenchmarkPopCountWithExprExpect32(b *testing.B) {
	benchmarkPopCountWithExpr(b, expect32)
}
func BenchmarkPopCountWithExprExpect64(b *testing.B) {
	benchmarkPopCountWithExpr(b, expect64)
}

func BenchmarkPopCountWithLoopExpect0(b *testing.B) {
	benchmarkPopCountWithLoop(b, 0)
}
func BenchmarkPopCountWithLoopExpect16(b *testing.B) {
	benchmarkPopCountWithLoop(b, expect16)
}
func BenchmarkPopCountWithLoopExpect32(b *testing.B) {
	benchmarkPopCountWithLoop(b, expect32)
}
func BenchmarkPopCountWithLoopExpect64(b *testing.B) {
	benchmarkPopCountWithLoop(b, expect64)
}

func BenchmarkPopCountWithBitwiseExpect0(b *testing.B) {
	benchmarkPopCountWithBitwise(b, 0)
}
func BenchmarkPopCountWithBitwiseExpect16(b *testing.B) {
	benchmarkPopCountWithBitwise(b, expect16)
}
func BenchmarkPopCountWithBitwiseExpect32(b *testing.B) {
	benchmarkPopCountWithBitwise(b, expect32)
}
func BenchmarkPopCountWithBitwiseExpect64(b *testing.B) {
	benchmarkPopCountWithBitwise(b, expect64)
}

func BenchmarkPopCountWithClearingLSBExpect0(b *testing.B) {
	benchmarkPopCountWithClearingLSB(b, 0)
}
func BenchmarkPopCountWithClearingLSBExpect16(b *testing.B) {
	benchmarkPopCountWithClearingLSB(b, expect16)
}
func BenchmarkPopCountWithClearingLSBExpect32(b *testing.B) {
	benchmarkPopCountWithClearingLSB(b, expect32)
}
func BenchmarkPopCountWithClearingLSBExpect64(b *testing.B) {
	benchmarkPopCountWithClearingLSB(b, expect64)
}
