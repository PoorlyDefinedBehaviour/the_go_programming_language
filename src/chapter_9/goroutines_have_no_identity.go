package main

// In most operating systems and programming languages that support multithreading,
// the current thread has a distinct identity that can be easily obtained as an ordinary value,
// typically an integer or pointer. This makes it easy to build an
// abstraction called thread-local storage, which is essentially a global map
// keyed by thread identity, so that each thread can store and retrieve
// values independent of other threads.
//
// Goroutines have no notion of identity that is acessible to the programmer.
// This is by design, since thread-local storage tends to be abused.
// For example, in a web server implemented in a language with thread-local storage,
// it's common for many functions to find information about the HTTP request
// on whose behalf they are currently working by in that storage.
// This is like acessing global variables which is considered harmful.
//
// Summary:
// In some languages, threads have identifiers called thread identity
// which programmers can access and build abstractions on top of, such as
// thread-local storage.
//
// Goroutines have no identity.

func main() {

}
