package main

import "fmt"

/*
Calling a function makes a copy of each argument value.
If a function needs to update a variable, or if an argument
is so large that we wishj to avoid copying it, we must pass
the address of the variable using a pointer. The same goes for
methods that need to update the receiver variable:
we attach them to the pointer type, sucih as *Point
*/

type Point struct {
	X float64
	Y float64
}

// Pointer type receiver.
// The name of this method is (*Point).scaleBy.
// The parentheses are necessary, without them, the expression
// would be parsed as *(Point.scaleBy)
func (point *Point) scaleBy(factor float64) {
	point.X *= factor
	point.Y *= factor
}

func (point Point) foo() string {
	return "foo"
}

// Named types (Point) and pointers to them (*Point) are
// the only types that may appear in a receiver declaration.
// Types where the underlying type is a pointer or an
// interface can't be receivers.
//
// type P *int
// func (p P) f() { /* ... */} // compile error: invalid receiver type

func main() {
	point1 := &Point{X: 1, Y: 2}

	point1.scaleBy(2)

	fmt.Println(point1)

	point2 := &Point{X: 3, Y: 4}

	// When we call a function has a pointer receiver
	// but the variable is a not a pointer, the compiler will
	// implicitly take the address of the variable using &point2
	// if the variavble is addressable
	point2.scaleBy(2) // same as (&point2).scaleBy(2)

	fmt.Println(point2)

	// Point{X:1, Y:2}.scaleBy(2) compile error: can't take address of Point literal
	// since point is a rvalue, it has no address and therefore the compiler
	// can't create a pointer to it.

	// When we have a pointer and call a method that has a non-pointer
	// receiver, the compiler will implictly deference the pointer for us.
	fmt.Println(point1.foo() == (*point1).foo()) // true
}
