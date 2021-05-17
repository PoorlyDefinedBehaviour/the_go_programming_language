package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
On occasion, a program must take different actions depending
on the kind of error that has occurred. Consider an attempt to read
n bytes of data from a file. It n is chosen to be the length of the file,
any error represents a failure. On the other hand, if the caller
repeatedly tries to read fixed-size chunks until the file is exhausted,
the caller must response differently to and-of-file condition than it does
to all other errors.
For this reason, the io package guarantees that any read failure by an end-of-file
condition is always reported by a distinguished error, io.EOF, which is defined as follows:

package io

// EOF is the error returned by Read when no more input is available
var EOF = errors.New("EOF")
*/

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF { // the caller can check for io.EOF using the == operator
			break

		}

		if err != nil {
			panic(err)
		}

		fmt.Println(r)
	}
}
