package main

import (
	"fmt"
	"runtime"
)

// GOOS and GOARCH can be set during compilation
//
// Example:
// GOARCH=386 go build main.go
//
// Some packages may need to compile differente versions of
// the code for certain platforms or processors,
// to deal with low-level portability issues or to provide
// optimized versions of important routines.
// If a file name includes an operating_system or processor
// architecture name like net_linux.go or asm_amd64.s, then
// the go tool will compile the file only when building for that target.
//
// Special comments called build tags give more fined-grained control.
// For example, if a file contains this comment:
// // +build linux darwin
// before the package declaration(and its doc comment),
// go build will compile it only when building for Linux or Mac OS X,
// and this comment saays never to compile the file:
// // +build ignore

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
