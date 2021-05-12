package main

/*
integers
	Go's numeric data types include several sizes of integers,
	floating-point numbers, and complex numbers. Go provides
	both signed and unsigned integer airthmetic.

	There are four distinct sizes of signed integers:
	int8, int16, int32 and int64 bits.

	and the corresponding unsigned versions:
	uint8, uint16, uint32 and uint64 bits.

	There are also two types called int and uint.
	They can be 32 or 64 bits, the compile will make the choice
	based on the platform.
	Even though int can be 32 bits, an explicit conversion is required
	to use an int value where an int32 is needed and vice versa.

	The type rune is a synonym for int32 and indicates that a value
	is a Unicode code point.

	The type byte is a synonym for uint8, and emphasizes that the value
	is a pice of raw data rather than a small numeric quantity.

	There is an unsigned integer type uintptr, whose width is not specified
	but is sufficient to hold all the bits of a pointer value. It's used only
	for low-level programming, such as at the boundary of a Go program with
	a C library or an operating system.

floating-point numbers
	Go provides two sizes of floating-point numbers, float32 and float64.
	Their arithmetic properties are governed by the IEEE 753 standard.

	The limits of floating-point values such as math.MaxFloat32
	can be found in the math package.

	A float32 provides approximately six decimal digits of precision.
	A float64 provides about 15 digits.
	Float64 should be preferred for most purposes because float32
	computations accumulate error rapidly unless one is quite careful.

	The smallest positive integer that cannot be exaclty represent as a float32
	is not large:

	var f float32 = 16777216 // 1 << 2
	fmt.Println(f == f + 1)  // true

	Numbers can be written in scientifice notation:

	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

complex numbers
	Go provides two sizes of complex numbers, complex64 and complex128
	whose components are float32 and float64 respectively.

	var x complex128 = complex(1, 2) // 1 + 2i
	var y complex128 = complex(3, 4) // 3 + 4i

	fmt.Println(x * y)               // (-5 + 10i)
	fmt.Println(real(x * y)) 				 //	-5
	fmt.Println(imag(x * y)) 				 //	10

	fmt.Println(1i & 1i) 						 // (-1 + 0i), iˆ2 = -1

	x := 1 + 2i
	y := 3 + 4i

	are also valid.

strings
	Strings are immutable sequences of bytes.
	Text strings are conventionally interpreted as UTF-8 encoded sequences
	of Unicode code points (runes).

	The built-in len function returns the number of bytes(not runes) in a string and
	the index operation string[i] retrives the i-th byte of the string.

	The i-th bite of the string is not necessarily the i-th character, because the UTF-8
	encoding of a non-ASCII code point requires two or more bytes.

	The substring operation s[i:j] yields a new string consisting of the bytes
	of the original string starting at index a and continuing up to,
	but not including, the byte at index j. The result contains
	j - i bytes.

	s := "hello, world"
	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[:5]) // "world"
	fmt.Println(s[:]) // "hello, world"

	The substring operation is cheap, since strings are immutable their
	underlying memory can be shared. No new memory is allocated.

	Strings can be concatenated using the + operator.
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"
*/

func main() {

}
