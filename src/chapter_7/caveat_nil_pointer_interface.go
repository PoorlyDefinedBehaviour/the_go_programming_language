package main

import (
	"bytes"
	"io"
)

func f(out io.Writer) {
	// This check fails because a nil pointer is passed
	// as an interface and the interface dynamic type is
	// *bytes.Buffer in this case
	if out == nil {
		return
	}

	// will panic
	out.Write([]byte("done\n"))
}

func main() {
	// Caveat: An interface containing a nil pointer is not nill
	// A nil interface value, which contains no value at all,
	// is not the same as an interface value containing a pointer
	// that happens to be nil.

	var buffer *bytes.Buffer

	f(buffer)
}
