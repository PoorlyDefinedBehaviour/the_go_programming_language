package main

import "fmt"

func main() {
	// Slices represent the variable length sequences whose elements
	// all have the same type.
	// A slice type is written []T, where elements have type T. It looks
	// like an array without a size
	//
	// A slice has an underlying array.
	//
	// A slice has 3 components: a pointer, a length and a capacity.
	// The pointer points to the first element of the array that is
	// reachable through the slice, which is not necessarily the array's first element
	//
	// The length is the number of slice elements. It can't exceed the capacity,
	// which is usually the number of elements between the start of the slice and the end
	// of the underlying array.
	//
	// len(slice) returns the length of the slice
	// cap(slice) returns the capcity of the slice which is the number
	// os elements between the start of the slice and the end of the underlying array.
	//
	// Multiple slices can share the same underlying array.

	months := [...]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)     // ["April", "May", "June"]
	fmt.Println(summer) // ["June", "July", "August"]

	// Slicing past cap(slice) causes a panic.
	// Slicing beyong len(slice) extends the slice.

	// fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extends a slice (within capacity)
	// works because summer is a slice of months which has a capacity
	// of at least 13

	fmt.Println(endlessSummer)

	// Slicing takes O(1) time
	// Since a slice contains a pointer to an element of an array,
	// passing a slice to a function permits the function to
	// modify the underlying array elements.

	// Slices can't be compared using the == operator.
	// The zero value of a slice type is nil. A nil slice has no underlying array
	// and has length and capacity zero, but there are also non-nill slices of length
	// and capacity zero.
	fmt.Println(len([]int{}))            // 0
	fmt.Println(cap(make([]int, 3)[3:])) // 0

	// As with any type that can have nil values, the nil value of a particular slice
	// type can be written using a conversion expression such as []int(nil).
	fmt.Println(len([]int(nil)))   // 0
	fmt.Println([]int(nil) == nil) // true
	fmt.Println([]int{} == nil)    // false

	// The built-in function make creates a slice of a specified element type, length and capacity.
	// The capacity argument may be omitted, in which case the capacity equals the length.
	// make([]T, len)
	// make([]T, len, cap) which is the same as make([]T, cap)[:len]
	// Under the hood, makes creates an unnamed array variable and returns a slice of it.
	// The array is accessible only through the return slice.
	// In the first form of make the slice is a view of the entire array.
	// In the second form of make the slice is a view of only the array's first len elements,
	// but its capacity includes the entire array.
}
