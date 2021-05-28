package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

/*
Conceptually, a value of an interface type, or interface value,
has two components, a concrete type and a value of that type.
These are called the interface's dynamic type and dynamic value.

How it actually looks like in the Go source code:

type iface struct {
  tab  *itab // points to type
  data unsafe.Pointer // points to value
}

// empty interface
type eface struct {
	_type *_type
	data  unsafe.Pointer
}

At compile time, Go doesn't know what the dynamic type of an interface
value will be, so a call through an interface must use dynamic dispatch.
Instead of a direct call, the compiler must generate code to obtain
the address of the method named Write from the type descriptor,
then make an indirect call to that address
*/

func main() {
	// Go variablesare always initialized to a well-defined value, and interfaces are no exception.
	// The zero value for an interface has both its type and value components set to nil.
	//
	// how writer looks like:
	// writer
	// type -> nil
	// value -> nil
	var writer io.Writer

	fmt.Println(writer == nil) // true because both type and value are nil

	// writer.Write([]byte("hello")) // panic: runtime error: invalid memory address or nil pointer dereference

	// how writer looks like:
	// writer
	// type -> *os.File
	// value -> os.File{ fd int = 1(stdout) }
	writer = os.Stdout // equivalent to writer = io.Writer(os.Stdout)

	// Calling Write on a interface value containing an *os.File
	// causes the (*os.File).Write method to be called.
	// The effect is as if we had made this call directly:
	// os.Stdout.Write([]byte("hello"))
	writer.Write([]byte("hello")) // hello

	// how writer looks like:
	// writer
	// type -> *bytes.Buffer
	// value -> bytes.Buffer { data []byte }
	writer = new(bytes.Buffer)

	// Calls (*bytes.Buffer).Write with the address of the buffer as the receiver
	writer.Write([]byte("hello")) // hello

	// this resets both its components to nil, restoring w to th same state as when it was declared.
	// how writer looks like:
	// type -> nil
	// value -> nil
	writer = nil

	// interface values can hold any value
	var x interface{} = time.Now()

	x = []int{1, 2, 3}

	// Interface values may be compared
	// if the value it holds is comparable.
	// When the value it holds it no comparable, the runtime will panic.
	// fmt.Println(x == x) // panic because []int is not comparable

	// We can print out the dynamic value of an interface using %T.
	// Internally, fmt uses reflection to obain the name of the
	// interface's dynamic type.
	fmt.Printf("\n%T\n", x) // []int
}
