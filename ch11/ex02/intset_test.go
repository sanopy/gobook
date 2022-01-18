package intset

import (
	"testing"
)

func TestHas(t *testing.T) {
	var x IntSet
	var _x MapIntSet
	x.Add(0)
	x.Add(64)
	_x.Add(0)
	_x.Add(64)
	t.Run("exist target element", func(t *testing.T) {
		want := _x.Has(64)
		if got := x.Has(64); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
	t.Run("not exist target element", func(t *testing.T) {
		want := _x.Has(63)
		if got := x.Has(63); got != want {
			t.Errorf("actual: %v, want: %v", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	var x IntSet
	var _x MapIntSet
	x.Add(0)
	x.Add(63)
	x.Add(64)
	x.Add(127)
	_x.Add(0)
	_x.Add(63)
	_x.Add(64)
	_x.Add(127)
	want := _x.String()
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	var _x, _y MapIntSet
	x.Add(0)
	x.Add(64)
	x.Add(127)
	y.Add(0)
	y.Add(63)
	y.Add(128)
	_x.Add(0)
	_x.Add(64)
	_x.Add(127)
	_y.Add(0)
	_y.Add(63)
	_y.Add(128)
	x.UnionWith(&y)
	_x.UnionWith(&_y)
	want := _x.String()
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}
func TestString(t *testing.T) {
	var x IntSet
	var _x MapIntSet
	x.Add(0)
	x.Add(127)
	x.Add(63)
	x.Add(64)
	x.Add(8)
	x.Add(4)
	_x.Add(0)
	_x.Add(127)
	_x.Add(63)
	_x.Add(64)
	_x.Add(8)
	_x.Add(4)
	want := _x.String()
	if got := x.String(); got != want {
		t.Errorf("actual: %v, want: %v", got, want)
	}
}
