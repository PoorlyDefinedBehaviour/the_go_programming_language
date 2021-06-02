package main

import (
	"bytes"
	"io"
	"os"
)

// A type assertion is an operation applied to an interface value.
// Syntactically, it looks like x.(T),
// where x is an expression of an interface type and T is a type,
// called the "asserted" type. A type ssertions checks that
// the dynamic type of its operand matches the asserted type.

// There are two possibilities. First, if the asserted type T is
// a concrete type, then the type assertions checks whether
// x's dynamic type is identical to T. If this check succeeds,
// the result of the type assertion is x's dynamic value, whose type
// is of course T. If the check fails, then the operation panics.
func example1() {
	var writer io.Writer = os.Stdout

	f := writer.(*os.File)      // success: f == os.Stdout
	c := writer.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
}

type ByteCounter int

// io.Writer implementation.
func (byteCounter *ByteCounter) Write(p []byte) (bytesWritten int, err error) {
	*byteCounter += ByteCounter(len(p))

	bytesWritten = len(p)
	err = nil

	return bytesWritten, err
}

// If the asserted type T is an interface type, then the type assertions
// checks whether x's dynamic type satisfier T. If this check succeeds,
// the dynamic value is not extracted. A type assertions to an interface
// type changes the type of the expression, making a different set
// of methods accessible.
func example2() {
	var writer io.Writer = os.Stdout

	readWriter := writer.(io.ReadWriter) // successs: *os.File has both Read and Write

	writer = new(ByteCounter)

	readWriter = writer.(io.ReadWriter) // panic: *ByteCounter has no Read method
}

// When the operand of a type assertion is nil, the operation will always fail.
func example3() {
	var writer io.Writer

	rw := writer.(io.ReadWriter) // fails
}

// We can check if the dynamic type of an interface value is some particular type.
func example4() {
	var writer io.Writer = os.Stdout

	f, ok := writer.(*os.File)      // ok == true
	b, ok := writer.(*bytes.Buffer) // ok == false
}

func main() {

}
