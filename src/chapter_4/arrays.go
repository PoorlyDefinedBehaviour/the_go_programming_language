package main

import "fmt"

func main() {
	// ARRAYS ARE PASSED TO FUNCTIONS BY COPY
	/*
		An array is a fixed-length sequence of zero or more elements of a particular type.

		Individual array elements are accsssed with the conventional subscript notation.

		The built-in function len returns the number of elements in the array.
	*/

	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	/*
		By default, the elements of a new array variable are initially set to the zero
		value of the type that the array holds.
	*/

	// We can use an array literal to initialize an array with a list of values
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q) // [1 2 3]

	// In an array literal, if an ellipsis "..." appears in place of the length,
	// the array length is determined by the number of initializers.
	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // [3]int

	/*
		The size of an array is part of its type, so [3]int and [4]int are differente types.
		The size must be a constant expression.
	*/

	// It is also possible to specify a list of index and value pairs.
	// In this form, indices, can appear in any order and some may be ommited
	// this line defines an array with 100 elements, all initialized to the zero value of int expect for the last,
	// which has value -1
	r := [...]int{99: -1}
	fmt.Println(r)      // [0 0 0 ... 0 -1]
	fmt.Println(len(r)) // 100

	// If an array's element type is comparable then the array type is comparable too.
	// We can compare two arrays of that type using the == operator.
	aa := [2]int{1, 2}
	bb := [...]int{1, 2}
	cc := [2]int{1, 3}

	fmt.Println(aa == bb, aa == cc, bb == cc) // true false false
}
