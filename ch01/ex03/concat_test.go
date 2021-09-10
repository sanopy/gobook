package main

import (
	"testing"
)

var arr = []string{"Austin", "Bexar", "Cactus", "Diablo", "Essex", "Folsom", "Grizzly",
	"Havana", "Icehouse", "Juno", "Kilo", "Liberty", "Mitaka", "Newton", "Ocata", "Pike",
	"Queens", "Rocky", "Stein", "Train", "Ussuri", "Victoria", "Wallaby", "Xena", "Yoga"}

func BenchmarkStringsPlusOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsPlusOperator(arr)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin(arr)
	}
}
