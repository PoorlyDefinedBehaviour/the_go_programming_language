package main

import (
	"fmt"
	"strings"
)

/*
Functions are first-class values in Go: like other values,
function values have types, and they may be assigned to variables
or passed to or returned from functions.

The zero value of a function type is nil.
Calling a nil function causes a panic.

var f func(int) int
f(3) // panic: call of nil function
*/

var f func(int) int

var g func(int) int

func add1(r rune) rune {
	return r + 1
}

func main() {
	// Function values may be compared with nil
	if f != nil {
		f(3)
	}

	// Functions are not comparable, and because of that
	// they can't be used as keys in a map.
	if f == g { // compiler error

	}

	// Since functions are values, we can have higher order functions
	// strings.Map(f, string) applies F to each character of the string
	fmt.Println(strings.Map(add1, "HAL-9000")) // IBM:.:1111
}
