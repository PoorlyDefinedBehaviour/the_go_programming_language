package main

import "fmt"

/*
A variadic function is one that can be called with varying numbers of arguments.
The most familia examples are fmt.Printf and its variants.
Printf requires one fixed argument at the beginning,
then acepts any number of subsequent arguments.

To declare a variadic function, the type of the final parameter is preceded
by an ellipsis "...", which indicates that the function may be called
with any number of arguments of this type.

Implicitly, the caller allocates an array, copies the arguments into it,
and passes a slice of the entire array to the function.
*/

/*
Although xs ...int behaves like a slice, its type is no the same type
as the type of a slice.
*/
func sum(xs ...int) int {
	total := 0

	for _, x := range xs {
		total += x
	}

	return total
}

func min(x int, y int, xs ...int) int {
	var result int = x

	if y < result {
		result = y
	}

	for _, value := range xs {
		if value < result {
			result = value
		}
	}

	return result
}

func main() {
	fmt.Println(sum())              // 0
	fmt.Println(sum(3))             // 3
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	// slices can be "expanded" to match variadic functions
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...)) // 15

	fmt.Println(min(1, 2))          // 1
	fmt.Println(min(1, 2, 3, 4, 5)) // 1
	fmt.Println(min(5, 4, 3, 2, 1)) //1
}
