package main

import "os"

// Creates a set of directories and later removes them
func foo(directories []string) {
	rmdirFunctions := make([]func(), 0, len(directories))

	for _, directory := range directories {
		/*
		This is necessary because the for loop introduces a new lexical block
		in which the variable directory is declared. All function values created
		by this loop "capture" and share the same variable, an addressable storage location,
		not is value at that particular moment.
		*/
		directoryCopy := directory

		os.MkdirAll(directory, 0755) // 0755 makes parent directories be created too

		rmdirFunctions = append(rmdirFunctions, func() { os.RemoveAll(directoryCopy) })
	}

	for _, f := range rmdirFunctions {
		f()
	}
}

func main() {

}
