package intset

import (
	"math/rand"
	"testing"
	"time"
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
	y.Add(512)
	y.Add(1024)
	_x.Add(0)
	_x.Add(64)
	_x.Add(127)
	_y.Add(0)
	_y.Add(63)
	_y.Add(128)
	_y.Add(512)
	_y.Add(1024)
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

var times = 100
var seed = time.Now().UTC().UnixNano()

func BenchmarkIntSetHas_1e3(b *testing.B) {
	benchmarkIntSetHas(b, 1e3)
}
func BenchmarkIntSetHas_1e5(b *testing.B) {
	benchmarkIntSetHas(b, 1e5)
}
func BenchmarkIntSetHas_1e7(b *testing.B) {
	benchmarkIntSetHas(b, 1e7)
}
func benchmarkIntSetHas(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	var set IntSet
	for i := 0; i < times; i++ {
		set.Add(rng.Intn(n))
	}
	for i := 0; i < b.N; i++ {
		set.Has(rng.Intn(n))
	}
}

func BenchmarkMapIntSetHas_1e3(b *testing.B) {
	benchmarkMapIntSetHas(b, 1e3)
}
func BenchmarkMapIntSetHas_1e5(b *testing.B) {
	benchmarkMapIntSetHas(b, 1e5)
}
func BenchmarkMapIntSetHas_1e7(b *testing.B) {
	benchmarkMapIntSetHas(b, 1e7)
}
func benchmarkMapIntSetHas(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	var set MapIntSet
	for i := 0; i < times; i++ {
		set.Add(rng.Intn(n))
	}
	for i := 0; i < b.N; i++ {
		set.Has(rng.Intn(n))
	}
}

func BenchmarkIntSetAdd_1e3(b *testing.B) {
	benchmarkIntSetAdd(b, 1e3)
}
func BenchmarkIntSetAdd_1e5(b *testing.B) {
	benchmarkIntSetAdd(b, 1e5)
}
func BenchmarkIntSetAdd_1e7(b *testing.B) {
	benchmarkIntSetAdd(b, 1e7)
}
func benchmarkIntSetAdd(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set IntSet
		for i := 0; i < times; i++ {
			set.Add(rng.Intn(n))
		}
	}
}

func BenchmarkMapIntSetAdd_1e3(b *testing.B) {
	benchmarkMapIntSetAdd(b, 1e3)
}
func BenchmarkMapIntSetAdd_1e5(b *testing.B) {
	benchmarkMapIntSetAdd(b, 1e5)
}
func BenchmarkMapIntSetAdd_1e7(b *testing.B) {
	benchmarkMapIntSetAdd(b, 1e7)
}
func benchmarkMapIntSetAdd(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		var set MapIntSet
		for i := 0; i < times; i++ {
			set.Add(rng.Intn(n))
		}
	}
}

func BenchmarkIntSetUnionWith_1e3(b *testing.B) {
	benchmarkIntSetUnionWith(b, 1e3)
}
func BenchmarkIntSetUnionWith_1e5(b *testing.B) {
	benchmarkIntSetUnionWith(b, 1e5)
}
func BenchmarkIntSetUnionWith_1e7(b *testing.B) {
	benchmarkIntSetUnionWith(b, 1e7)
}
func benchmarkIntSetUnionWith(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	var s1, s2 IntSet
	for i := 0; i < times; i++ {
		s1.Add(rng.Intn(n))
		s2.Add(rng.Intn(n))
	}
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}

func BenchmarkMapIntSetUnionWith_1e3(b *testing.B) {
	benchmarkMapIntSetUnionWith(b, 1e3)
}
func BenchmarkMapIntSetUnionWith_1e5(b *testing.B) {
	benchmarkMapIntSetUnionWith(b, 1e5)
}
func BenchmarkMapIntSetUnionWith_1e7(b *testing.B) {
	benchmarkMapIntSetUnionWith(b, 1e7)
}
func benchmarkMapIntSetUnionWith(b *testing.B, n int) {
	rng := rand.New(rand.NewSource(seed))
	var s1, s2 MapIntSet
	for i := 0; i < times; i++ {
		s1.Add(rng.Intn(n))
		s2.Add(rng.Intn(n))
	}
	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}
