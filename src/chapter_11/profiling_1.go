package main

// Go supports many kinds of profiling, each concerned
// with a different aspect of performance,
// but all of them involve recording a sequence of events
// of interest, each of which has accompaying stack trace.
//
// A CPU profile identifies the functions whose execution
// requires the most CPU time. The currently running thread
// on each CPU is interrupted periodically by the operating system
// every fwe milliseconds, with each interruption recording one profile
// event before normal execution resumes.
//
// A heap profile identifies the statements responsible for allocating
// the most memory. The profiling library samples calls to
// the internal memory allocation routines so that on average,
// one profile event is recorded per 512KB of allocated memory.
//
// A blocking profile identifies the operations responsible for blocking
// goroutines the longest, such as system calls, channel sends and receives,
// and acquisitions of locks. The profiling library records an event
// every time a goroutine is blocked by one of these operations.
//
// How to profile:
// go test -cpiprofile=cpu.out
// go test -blockprofile=block.out
// go test -memprofile=mem.out
//
// Use one at a time because the instrumentation
// of one profiler may skew the results of other profiler.
//
// Profiling can be enabled programatically using the runtime API.
//
// Profiles can be analyzed using the pprof tool. This is a standard
// part of the Go distribution and it can be accessed using go tool pprof.
// To make profiling efficient and save space, the log does not include
// function names; instead, functions are identified by their addresses.
// This means that pprof needs the executable in order to make sense of the log.
// Although go test usually discards the test executable once the test
// is complete, when profiling is enabled it saved the executable as foo.test,
// where foo is the name of the tested package.

func main() {

}
