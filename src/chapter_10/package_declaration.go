package main

// A package declaration is required at the start of every Go source file.
// Its main purpise is to determine the default identifier for that package(called the package name)
// when it is imported by another package.
//
// Packages that are meant to be executed will always have the name main,
// regardless of the package's import path. This is a signal to go build that it
// must invoke the linker to make an executable file.
//
// Some files in the directory may have the suffix _test
// on their package name if the file name ends with _test.go.
// Such a directory may define two packages:
// the usual one, plus another one called an external tet package.
// The _test suffix signals to go test that it must build both packages,
// and it indicates which files belong to each package.
// External test packages are used to avoid cycles in the import graph
// arising from dependencies of the test.

func main() {

}
