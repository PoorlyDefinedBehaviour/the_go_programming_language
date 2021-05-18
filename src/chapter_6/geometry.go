package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

// Function.
func distance(from, to Point) float64 {
	return math.Hypot(to.X-from.X, to.Y-from.Y)
}

// Method.
// The extra parameter point Point is called the metho's receiver,
// a legacy from early object-oriented languages that describe
// calling a method as "sending a message to an object".
func (point Point) distance(to Point) float64 {
	return math.Hypot(to.X-point.X, to.Y-point.Y)
}

// Path is a named slice type, not a struct like Point, yts we can
// still define methods for it.
// Go allows any named type defined in the same package to have methods
// as long as the underlying type is neither a pointer nor an interface.
type Path []Point

// Since each type has its own name space for methods, we can use the
// same property/method name for different types.
func (path Path) distance() float64 {
	sum := 0.0

	for i := range path {
		if i == 0 {
			continue
		}

		sum += path[i-1].distance(path[i])
	}

	return sum
}

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}

	fmt.Println(distance(p, q)) // function call

	// In a method call, the receiver argument appears before the method name
	fmt.Println(p.distance(q)) // method call

	// The expression p.distance is called a selector, because it selects
	// the appropriate distance method for the receiver p of type Point.
	// Selectors are also used to select fields of struct types,
	// as in p.X. Since methods and fields inhabit the same name space,
	// declaring a method X on the struct type Point would be ambiguous
	// and the compiler will reject it.

	// ---
	perimeter := Path{{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}
	fmt.Println(perimeter.distance()) // 12
}
