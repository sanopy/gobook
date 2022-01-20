// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"sort"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type MapIntSet struct {
	elements map[int]bool
}

// Has reports whether the set contains the non-negative value x.
func (s *MapIntSet) Has(x int) bool {
	if s.elements == nil {
		return false
	}
	return s.elements[x]
}

// Add adds the non-negative value x to the set.
func (s *MapIntSet) Add(x int) {
	if s.elements == nil {
		s.elements = make(map[int]bool)
	}
	s.elements[x] = true
}

// UnionWith sets s to the union of s and t.
func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for x := range t.elements {
		s.Add(x)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	var num []int
	for x := range s.elements {
		num = append(num, x)
	}
	sort.Ints(num)

	for i, x := range num {
		if i == 0 {
			fmt.Fprintf(&buf, "%d", x)
		} else {
			fmt.Fprintf(&buf, " %d", x)
		}
	}

	buf.WriteByte('}')
	return buf.String()
}
