package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	// We need to keep track of the type because different
	// variables may have the same address.
	// Example:
	// x and y are both arrays, x and x[0] have the same address,
	// as do y and y[0].
	t reflect.Type
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	// A lot of checks were omitted for brevity.
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	if x.CanAddr() && y.CanAddr() {
		xpointer := unsafe.Pointer(x.UnsafeAddr())
		ypointer := unsafe.Pointer(y.UnsafeAddr())

		if xpointer == ypointer {
			return true
		}

		c := comparison{x: xpointer, y: ypointer, t: x.Type()}
		if seen[c] {
			return true
		}

		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Int:
		return x.Int() == y.Int()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}

		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}

		return true
	}

	panic("unreachable")
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func main() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
}
