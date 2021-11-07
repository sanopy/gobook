package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T) {
	var tr *tree
	tr = add(tr, 4)
	tr = add(tr, 2)
	tr = add(tr, 1)
	tr = add(tr, 3)
	tr = add(tr, 6)
	tr = add(tr, 5)
	tr = add(tr, 7)
	want := "(4 (2 (1) (3)) (6 (5) (7)))"
	if got := tr.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}
