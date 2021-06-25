package main

// Each OS thread has fixed-size block of memory(often as large as 2MB)
// for its stack. A 2MB stack would be a huge waste of memory
// for a little goroutine, such as one that merely waits for a
// WaitGroup then closes a channel. It's not uncommong for a GO
// program to create hundreds of thousands of goroutines at one tine,
// which would be impossible with stacks this large. Yet despite their size,
// fixed-size stacks are not always big enough for the most complex
// and deeply recursive of functions. Changing the fixed size can
// improve space efficiency and allow more threads to be created,
// or it can enable more deeply recursive functions, but it cannot do both.
//
// In constrast, a goroutine starts life with a small stack, typically 2KB.
// A goroutine's stack, like the stack of an OS thread,
// holds the local variables of active and suspended function calls,
// but unlike an OS thread, a goroutine's stack is NOT fixed;
// it grows and shrinks as needed. The size limit for a goroutine stack may
// be as much as 1GB, in practice, few goroutines use that much.

func main() {

}
