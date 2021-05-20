package main

import (
	"fmt"
	"time"
)

/*
Usually we select and call a method in the same expression,
as in point.distance(), but it's possible to separate
these two operations. The selector point.istance yields
a method value, a function that binds a method (Point.distance)
to a specific receiver value point.
*/
type Point struct {
	X, Y float64
}

func (point Point) distance(to Point) float64 {
	return 0.0
}

func (point *Point) scaleBy(factor float64) {
	point.X *= factor
	point.Y *= factor
}

func (point Point) add(other Point) Point {
	return Point{X: point.X + other.X, Y: point.Y + other.Y}
}

func (point Point) sub(other Point) Point {
	return Point{X: point.X - other.X, Y: point.Y - other.Y}
}

type Path []Point

// making []Point a functor would be better
// but since slices aren't functors,
// a custom package slices.Map(slice, f) would still be better
func (path Path) translateBy(offset Point, add bool) {

	var op func(p, q Point) Point

	if add {
		op = Point.add
	} else {
		op = Point.sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

type Rocket struct {
}

func (rocket *Rocket) launch() {}

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 3, Y: 4}

	distanceFromP := p.distance

	fmt.Println(distanceFromP(q))

	rocket := Rocket{}

	// since we can bind a function to a receiver,
	// point-free style is possible.
	time.AfterFunc(10*time.Second, rocket.launch)

	// Method expression: we can take a method and supply
	// the receiver at the call site.
	distance := Point.distance

	fmt.Println(distance(p, q)) // p is the receiver

	scale := (*Point).scaleBy

	scale(&p, 2)

	fmt.Println(p)
}
