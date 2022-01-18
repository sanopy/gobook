package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		input   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			input:   "",
			counts:  map[rune]int{},
			utflen:  [utf8.UTFMax + 1]int{},
			invalid: 0,
		},
		{
			input: "Hello, World!",
			counts: map[rune]int{
				'H': 1,
				'e': 1,
				'l': 3,
				'o': 2,
				',': 1,
				' ': 1,
				'W': 1,
				'r': 1,
				'd': 1,
				'!': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{0, 13, 0, 0, 0},
			invalid: 0,
		},
		{
			input: "日本語のテスト",
			counts: map[rune]int{
				'日': 1,
				'本': 1,
				'語': 1,
				'の': 1,
				'テ': 1,
				'ス': 1,
				'ト': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 7, 0},
			invalid: 0,
		},
		{
			input: "👋🤚👌🙏",
			counts: map[rune]int{
				'👋': 1,
				'🤚': 1,
				'👌': 1,
				'🙏': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 0, 4},
			invalid: 0,
		},
		{
			input: "Hello, 世界🤚",
			counts: map[rune]int{
				'H': 1,
				'e': 1,
				'l': 2,
				'o': 1,
				',': 1,
				' ': 1,
				'世': 1,
				'界': 1,
				'🤚': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{0, 7, 0, 2, 1},
			invalid: 0,
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		counts, utflen, invalid, err := charcount(r)
		if err != nil {
			t.Errorf("charcount: %v", err)
		}

		if !reflect.DeepEqual(counts, tt.counts) {
			t.Errorf("counts(%q) = %v, want %v", tt.input, counts, tt.counts)
		}

		if !reflect.DeepEqual(utflen, tt.utflen) {
			t.Errorf("counts(%q) = %v, want %v", tt.input, utflen, tt.utflen)
		}

		if invalid != tt.invalid {
			t.Errorf("counts(%q) = %q, want %q", tt.input, invalid, tt.invalid)
		}
	}
}
