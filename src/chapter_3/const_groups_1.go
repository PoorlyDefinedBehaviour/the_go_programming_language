package main

import "fmt"

const (
	a = 1
	b
	c = 2
	d
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(a, b, c, d) // 1 1 2 2
}
