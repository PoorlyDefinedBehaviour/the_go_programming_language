package main

import (
	"fmt"
	"os"
	"reflect"
)

// A variable is an addressable storage location that contains a value,
// and its value may be updated through that address.
//
// Some reflect.Value are addressable, others are not.

func main() {
	// No reflect.Value returned by reflect.ValueOf is addressable.
	x := 2                   // value type variable(addressable)?
	a := reflect.ValueOf(2)  // 2     int  no
	b := reflect.ValueOf(x)  // 2     int  no
	c := reflect.ValueOf(&x) // &x    *int no
	d := c.Elem()            // 2     int  yes
	// We can obtain an addressble reflect.Value for any variable x
	// by using reflect.ValueOf(&x).Elem()

	// We can ask reflect.Value wether it is addressable through its CanAddr method:
	fmt.Println(a.CanAddr()) // false
	fmt.Println(b.CanAddr()) // false
	fmt.Println(c.CanAddr()) // false
	fmt.Println(d.CanAddr()) // true

	// We obtain an addressable reflect.Value whenever we indirect through a pointer,
	// even if we started from a non-addressable value.

	// To recover the variable from an addressable reflect.Value:
	x = 2
	d = reflect.ValueOf(&x).Elem()
	// Addr() returns a reflect.Value holding a pointer to the variable.
	// Interface() returns an interface{} containing the pointer to the variable.
	// We type assert the interface to the type of the varaible.
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x) // 3

	// We can also update the variable referred to by an addressable reflect.Value
	// by calling reflect.Value.Set:
	// This call will panic if types don't match or
	// when reflect.Value is not addressable.
	d.Set(reflect.ValueOf(4))
	fmt.Println(x) // 4

	// We can access private fields through reflection but we can not modify them.
	// To check if a value is addressable and settable, we can use reflect.Value.CanSet()
	stdout := reflect.ValueOf(os.Stdout).Elem() // reflect.Value(*os.Stdout)
	fmt.Println(stdout.Type())                  // os.File

	fd := stdout.FieldByName("fd")

	// always false because fd is a private field
	if fd.CanSet() {
		fd.SetInt(2) // panic
	}
}
