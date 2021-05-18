package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X float64
	Y float64
}

func (point Point) scaleBy(scale float64) Point {
	return Point{X: point.X * scale, Y: point.Y * scale}
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// A ColoredPointer is not a Point, but it has a Point.
// The embedded field instructs the compiler to generate
// additional wrapper methods that delegate to the declared methods,
// equivalent to these:
//
// func (point ColoredPoint) scaleBy(scale float64) Point {
// 	return point.Point.scaleBy(factory)
// }

func main() {
	point := ColoredPoint{
		Point: Point{
			X: 1,
			Y: 2,
		},
		Color: color.RGBA{0, 0, 0, 0},
	}

	fmt.Println(point.X, point.Point.X)
}
