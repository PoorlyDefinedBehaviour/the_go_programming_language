package main

import (
	"errors"
	"fmt"
	"syscall"
)

// How os package does it:

type PathError struct {
	Op   string
	Path string
	Err  error
}

var ErrNotExist = errors.New("file does not exist")

func (err *PathError) Error() string {
	return fmt.Sprintf("%s %s: %s", err.Op, err.Path, err.Err)
}

func IsNotExist(err error) bool {
	var actualErr error = nil

	if pathError, ok := err.(*PathError); ok {
		actualErr = pathError.Err
	}

	return actualErr == syscall.ENOENT || actualErr == ErrNotExist
}

func main() {

}
