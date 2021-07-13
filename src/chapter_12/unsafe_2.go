package main

import (
	"fmt"
	"unsafe"
)

// Most pointer types are written *T, meaning a pointer to
// a variable of type T. The unsafe.Pointer type is a special
// kind of pointer that can hold the address of any variable.
// Of course, we can't indirect through an unsafe.Pointer
// using *p because we don't know what type that expression should have.
// Like ordinary pointers, unsafe.Pointers are comparable and may
// be compared with nil, which is the zero value of the type.
//
// An ordinary *T pointer may be converted to an unsafe.Pointer,
// and an unsafe.Pointer may be converted back to an ordinary pointer,
// not necessarily of the same type *T.
//
// In summary, an unsafe.Pointer lets us write arbitrary
// values to memory and thus subvert the type system.
// An unsafe.Pointer may also be converted to a uintptr
// that holds the pointer's numeric value, letting us
// perform arithmetic on addresses.

func Float64bits(float float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&float))
}

func example1() {
	var x struct {
		a bool
		b int16
		c []int
	}

	pointer := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	*pointer = 42

	fmt.Println(x.b) // 42
}

func example2() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// This code is broken.
	// Some garbage collectors move variables around in memory
	// to reduce fragmentation or bookkeeping. Garbage collectors
	// of this kind are known as moving GCS. When a variable is moved,
	// all pointers that hold the address of the hold location must be
	// updated to point to the new one. From the perspective of the garbage collector,
	// an unsafe.Pointer is a pointer and thus its value must change as variable moves, but
	// a uintptr is just a number so its value must not change. This code hides a pointer
	// from the garbage collector in the non-pointer variable temp.
	// By the time the second statement executes, the variable x could have moved
	// and the number in temp would no longer be the address of &x.b.
	temp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pointer := (*int16)(unsafe.Pointer(temp))
	*pointer = 42

	// NOTE: We fix it by making temp a pointer just like example1()
	pointer = (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pointer = 42
}

func example3() {
	// This code is broken.
	// There are no pointers that refer to the variable created
	// by new, so the garbage collector is entitle
	// to recycle its storage when this statement completes.
	// pointer is just a number because of the uintptr cast.
	pointer := uintptr(unsafe.Pointer(new(int)))

	*((*int)(unsafe.Pointer(pointer))) = 10
}

func main() {
	fmt.Printf("%#016x\n", Float64bits(1.0))

	example1()
}
