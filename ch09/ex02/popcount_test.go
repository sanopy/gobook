package popcount

import (
	"testing"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

var input uint64 = 0x1234567890ABCDEF
var output int

func BenchmarkPopCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCount(input)
	}
	output = s
}

func BenchmarkBitCount(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += BitCount(input)
	}
	output = s
}

func BenchmarkPopCountByClearing(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountByClearing(input)
	}
	output = s
}

func BenchmarkPopCountByShifting(b *testing.B) {
	var s int
	for i := 0; i < b.N; i++ {
		s += PopCountByShifting(input)
	}
	output = s
}
