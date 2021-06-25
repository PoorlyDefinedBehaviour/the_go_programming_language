package main

// The Go scheduler uses a parameter called GOMAXPROCS to determine
// how many OS threads may be actively executing Go code simultaneously.
// Its default value if the number of CPUs on the machine,
// so on a machine with 8 CPUs, the scheduler will schedule Go code
// on up to 8 OS threads at once(GOMAXPROCS is the n in m:n scheduling).
// Goroutines that are sleeping or blocked in a communication do not need a thread at all.
// Goroutines that are blocked in I/O or other system calls or are calling non-go functions,
// do need an OS threads, but GOMAXPROCS need not account for them.
//
// You can explictly control this parameter using GOMAXPROCS
// environment variable or the runtime.GOMAXPROCS function.
//
// Example:
// GOMAXPROCS=1 go run main.go
// GOMAXPROCS=2 go run main.go

func main() {

}
