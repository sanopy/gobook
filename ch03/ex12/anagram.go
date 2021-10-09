package anagram

import (
	"bytes"
	"sort"
)

type Bytes []byte

func (b Bytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Bytes) Len() int {
	return len(b)
}

func isAnagram(s1, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)

	sort.Sort(Bytes(b1))
	sort.Sort(Bytes(b2))

	return bytes.Equal(b1, b2)
}
