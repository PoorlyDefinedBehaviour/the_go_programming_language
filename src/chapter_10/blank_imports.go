package main

// It is an error to import a package into a file but not refer to the name
// it defines within that file.
// However, on occasion we must import a package merely for the side effects of doing so:
// evaluation of the initializer expressions of its package-level variables and
// exceution of its init function. To supress unused import error we would other encounter,
// we must use a renaming imort in which the alternative name is _, the blank identifier.
// As usual the blank identifier can never be referenced.

import _ "image/png" // register PNG decoder

func main() {

}
