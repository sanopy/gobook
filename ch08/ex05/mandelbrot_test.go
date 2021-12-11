package main

import (
	"io"
	"testing"
)

func BenchmarkGenerateImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateImage(io.Discard)
	}
}

func BenchmarkGenerateImageParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateImageParallel(io.Discard)
	}
}
