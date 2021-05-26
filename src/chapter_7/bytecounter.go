package main

import "fmt"

type ByteCounter int

// io.Writer implementation.
func (byteCounter *ByteCounter) Write(p []byte) (bytesWritten int, err error) {
	*byteCounter += ByteCounter(len(p))

	bytesWritten = len(p)
	err = nil

	return bytesWritten, err
}

func main() {
	var byteCounter ByteCounter

	byteCounter.Write([]byte("hello"))

	fmt.Println(byteCounter) // 5 because len([]byte("hello")) == 5

	byteCounter = 0

	name := "Dolly"

	// ByteCounter can be passed to fmt.Fprintf because it expects
	// an io.Writer which is implemented by ByteCounter
	fmt.Fprintf(&byteCounter, "hello, %s", name)

	fmt.Println(byteCounter) // 12 because len("hello, Dolly") == 12
}
