package cycle

import (
	"bytes"
	"testing"
)

func TestCycle(t *testing.T) {
	one := 1

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1 := make(chan int)

	type link struct {
		value string
		tail  *link
	}
	cyclicLink, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	cyclicLink.tail, b.tail, c.tail = b, c, cyclicLink

	type mystring string

	var iface1 interface{} = &one

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// basic types
		{1, false},
		{1.0, false},
		{"foo", false},
		{mystring("foo"), false},
		{complex(1, 1), false},
		// slices
		{[]string{"foo"}, false},
		{cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		// pointers
		{&one, false},
		{new(bytes.Buffer), false},
		{cyclePtr1, true},
		// functions
		{func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		// structs
		{cyclicLink, true},
		// interfaces
		{&iface1, false},
	} {
		if IsCycle(test.x) != test.want {
			t.Errorf("IsCycle(%v) = %t",
				test.x, !test.want)
		}
	}
}
