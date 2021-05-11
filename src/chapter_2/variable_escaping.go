package main

var global *int

// x must be heap-allocated because it is still reachable
// from the variable global after f() has returned despite being
// declared as a local variable.
// we say that that x escapes f
func f() {
	var x int
	x = 1
	global = &x
}

// even though y is heap-allocated,
// it can freed after g() returns
// because y does not escape g()
func g() {
	y := new(int) // new(T) returns a pointer to T initialized to the zero value of T
	*y = 1
}

func main() {

}
