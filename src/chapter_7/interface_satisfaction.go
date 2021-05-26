package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
A type satisfies an interface if it possesses all the methods the interface requires.
For example, an *os.File satifies io.Reader, io.Writer, io.Closer and io.ReadWriter.
A *bytes.Buffer satisfiers io.Reader, io.Writer, and io.ReadWriter, but does not
satisfy io.Closer because it does not have a Close method.
*/

func main() {
	// Assignability: an expression may be assigned to an interface
	// only if its type satisfies the interface.

	var writer io.Writer

	writer = os.Stdout // OK: *os.File has Write method

	writer = new(bytes.Buffer) // OK: *bytes.Buffer has Write method

	// writer = time.Second // compile error: time.Duration lacks Write method

	fmt.Println(writer)
}
