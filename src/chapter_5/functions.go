package main

import (
	"fmt"
	"math"
)

/*
A function declaration has a name, a list of parameters, an optional list of results, and a body:
func name(parameter-list?) (result-list)? {
	body
}

A function can return no results, an unnamed result and one or more named results.

Leaving off the result list entirely implies that the function is not pure and
it's being called only for it's side effects.
---
*/
func hypot(x float64, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

/*
Functions arguments are passed by value, so modifications to the arguments
are local to the function unless the argument is a
pointer, slice, map, function or channel.
*/
func f(xs []int) {
	// this is a side effect
	xs[0] = 5
}

/*
When a function is declared without a body, it is an indication
that the function is implemented in a language other than Go.
Such declaration defines the function signature.

example:

package math

func Sin(x float64) float64 // implemented in assembly language
*/

func main() {
	fmt.Println(hypot(3, 4))

	xs := []int{1, 2, 3}
	f(xs)
	fmt.Println(xs) // [5 2 3]
}
