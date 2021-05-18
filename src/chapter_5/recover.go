package main

import "fmt"

/*
If the built-in recover function is called within a deferred function
and the function containing the defer statement is panicking,
recover ends the current state of panic and returns the value passed to panic().
The function that was panicking does not continue where it left off but
returns normally. If recover is called at any other time, it has no effect
and returns nil.
*/

type Syntax struct{}

func Parse(input string) (syntax *Syntax, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = fmt.Errorf("internal error: %v", reason)
		}
	}()

	// do some parsing ...
	return syntax, err
}

// Exercise 5.19
// Use panic and recover to write a function that contains no
// return statement yet returns a non-zero value
func exercise519() (result int) {
	defer func() {
		result = recover().(int)
	}()

	panic(1)
}

func main() {
	fmt.Println(exercise519()) // 1
}
