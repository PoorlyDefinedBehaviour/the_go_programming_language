package main

// Go compilation is notably faster than most other compiled languages,
// even when building from scratch. There are three main reasons
// for the compiler's speed. First, all imports must be explicitly
// listed at the beginnng of each source file, so the compiler does not
// have to read and process an entire file to determine its dependencies.
// Second, the dependencies of a package form a directed acyclic graph,
// and because there are no cycles, packages can be compiled separately and parhaps in parallel.
// Finally, the object file for a compiled Go package records
// export information not just for the package itself, but for its
// dependencies too.
// When compiling a package, the compiler must read one object file for each import
// but need not look beyond these files.
//
// Summary:
// Imports must be listed at the beginning of the file.
// Imports form a directed acyclic graph.
// Object file exports information for the package itself and for its dependencies.

func main() {

}
