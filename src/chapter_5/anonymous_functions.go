package main

import (
	"fmt"
	"strings"
)

/*
Named functions can be declared only at the package level,
but we can use a function literal to denote a function value
within any expression. A function literal is an expression and
is also called an anonymous function.
*/

func makeSquares() func() int {
	x := 0

	return func() int {
		x++

		return x * x
	}
}

func main() {
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))

	var i rune = 1

	// anonymous functions close over the environment
	fmt.Println(strings.Map(func(r rune) rune { return r + i }, "HAL-9000"))

	squares := makeSquares()

	fmt.Println(squares()) // 1
	fmt.Println(squares()) // 4
	fmt.Println(squares()) // 9
	fmt.Println(squares()) // 16
	fmt.Println(squares()) // 25
}
