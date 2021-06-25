package main

// OS threads are schedule by the OS kernel.
// Every few milliseconds, a hardware timer interrupts the processor,
// which causes a kernel function called the scheduler to be invoked.
// This function suspends the currently executing thread and saves its
// register in memory, looks over the list of threads and decides
// which one should run next, restores that thread's registers from memory,
// then resumes the execution of that thread.
// Because OS threads are scheduled by the kernel, passing control from
// one thread to another requires a full context switch, that is,
// saving the state of one user thread to memory, restoring the state of another,
// and updating the scheduler's data structures.
// This operation is slow, due to its poor locality and the number
// of memory accesses required, and has historically only gotten worse
// as the number of CPU cycles required to access memory has increased.
//
// The Go runtime contains its own scheduler that uses a technique known as
// m:n scheduling, because it multiplexes(or schedules) m goroutines on n OS threads.
// The job of the Go scheduler is analogous to that of the kernel scheduler,
// but it is concerned only with the goroutines of a single Go program.
//
// Unlike the operating system's thread scheduler, the Go scheduler is not invoked
// periodically by a hardware timer, but implicitly by certain Go language constructs.
// For example, when a goroutine calls time.Sleep or blocks in a channel or mutex operation,
// the scheduler puts it to sleep and runs another goroutine until it is time to wake the first one up.
// Because it doesn't need a switch to kernel context,
// rescheduling a goroutine is much cheaper than rescheduling a thread.
//
// Summary:
// Go schedules m goroutines on n OS threads.
// It is fast because it does not required a full context switch every time a reschedule happens.
// Rescheduling is triggered by language constructs, like time.Sleep or blocking in a channel.

func main() {

}
