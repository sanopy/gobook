package intset

import (
	"testing"
)

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(64)
	t.Run("exist target element", func(t *testing.T) {
		want := true
		if got := x.Has(64); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
	t.Run("not exist target element", func(t *testing.T) {
		want := false
		if got := x.Has(63); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	want := "{0 63 64 127}"
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(0, 63, 64, 127)
	want := "{0 63 64 127}"
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(0)
	x.Add(64)
	x.Add(127)
	y.Add(0)
	y.Add(63)
	y.Add(128)
	x.UnionWith(&y)
	want := "{0 63 64 127 128}"
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	want := 4
	if got := x.Len(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	t.Run("not exist target element", func(t *testing.T) {
		want := "{0 63 64 127}"
		x.Remove(1)
		x.Remove(128)
		if got := x.String(); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
	t.Run("exist target element", func(t *testing.T) {
		want := "{0 63 127}"
		x.Remove(64)
		if got := x.String(); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	x.Clear()
	want := "{}"
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	y := x.Copy()
	want := "{0 63 64 127}"
	if got := y.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}
