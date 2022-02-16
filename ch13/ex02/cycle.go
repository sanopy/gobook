package cycle

import (
	"reflect"
	"unsafe"
)

func isCycle(x reflect.Value, seen map[memo]bool) bool {
	if !x.IsValid() {
		return false
	}

	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := memo{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
		defer func() { seen[c] = false }()
	}

	switch x.Kind() {
	case reflect.Bool, reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return false

	case reflect.Ptr, reflect.Interface:
		return isCycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	panic("unreachable")
}

func IsCycle(x interface{}) bool {
	seen := make(map[memo]bool)
	return isCycle(reflect.ValueOf(x), seen)
}

type memo struct {
	x unsafe.Pointer
	t reflect.Type
}
