package main

// Although the unsafe package appears to be a regular package and is
// imported in the usual way, it is actually implemented by the compiler.
// It provides access to a number of built-in language features that are not
// ordinarily available because they expose details of Go's memory layout.
import (
	"fmt"
	"unsafe"
)

func main() {
	// unsafe.SizeOf(value) reports the size in bytes of the representation
	// of its argument which may be an expression of any type.
	// The expression is not evaluated.
	fmt.Println(unsafe.Sizeof(0))

	// A call to SizeOf is a constant expression of type uintptr,
	// so the result may be used as the dimension of an array type,
	// or to compute other constants.
	var xs [unsafe.Sizeof(0)]int
	fmt.Println(xs)

	// Sizeof reports only the size of the fixed part of each data structure,
	// like the pointer and length of a string but not indirect parts
	// like the contents of the string.
	//
	// Computers load and store values from memory most efficiently when those
	// values are properly aligned. For example, the address of a value of a two-byte
	// type such as int16 should be an even number, the address of a four-byte
	// value such as a rune should be a multiple of four, and the address
	// of an eight-byte value such as float64, uint64 or 64-bit pointer
	// should be a multiple of eight. Alignment requirements of higher
	// multiples are unusual, even for larger data types such as complex128.
	//
	// For this reason, the size of an aggregate type(a struct or array)
	// is at least the sum of the sizes of its fields or elements but may be
	// greater due to the presence of unused spaces added by the compiler
	// to ensure that the following field or element is propertly aligned
	// relative to the start of the struct or array.
	//
	// Average sizes:
	// Type														Size
	// bool 													1 byte
	// intN, uintN, floatN, complexN	N/8 bytes (for example, float64 is 8 bytes)
	// int, uint, uintptr							1 word(a word is usually 4 bytes)
	// *T															1 word
	// string													2 words(data, len)
	// []T														3 words(data, len, cap)
	// map														1 word
	// func														1 word
	// chan 													1 word
	// interface											2 words(type, value)
	//
	// The language specification does not guarantee that the order in which fields
	// are declared is the order in which they are lait out in memory,
	// so in theory a compiler is free to rearrange them,
	// although as we write this, none do. If the types of a struct's fields
	// are of different sizes, it may be more space-efficient to declare
	// the fields in an order that packs them as tightly as possible.
	//
	// The three structs below have the same fields, but the first requeis up to 50%
	// more memory than the other two:
	//																	64-bit 	32-bit
	// struct { bool; float64; int16} 	3 words	4 words
	// struct { float64; int16; bool }	2 words 3 words
	// struct { bool; int16; float64}		2 words 3 words
}
