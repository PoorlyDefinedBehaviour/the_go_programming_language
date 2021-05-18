package main

/*
Go's type system catches many mistakes at compile time, but others,
like an out-of-bounds array access or nil pointer deferefence,
require checks at run time. When the Go runtime detects
these mistakes, it panics.

During a typical panic, normal execution stops, all deferred
function calls in that goroutine are execute, and the program
crashes with a log message.

Not all panics come from the runtime. The built-in panic function
may be called directly; it accepts any value as an argument.

A panic is often called when some impossible situation happens,
which is often the indication of a bug.
*/
func main() {

}
