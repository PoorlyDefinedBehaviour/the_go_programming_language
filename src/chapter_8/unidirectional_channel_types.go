package main

import "fmt"

/*
The Go type system provides unidirectional channel types
that expose only one or the other of the send and receive operations.
The type chan<- int, a send-only channel of int, allows sends but not receives.
The type <-chan int, a receive-only channel of int, allows receives but not sends.

Calling close(channel) on a receive-only channel is a compile time error.
*/

func counter(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}

	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}

	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	// bidirectional channels
	naturals := make(chan int)
	squares := make(chan int)

	// counter takes an unidirectional send-only channel.
	// This is fine because implicit conversions from
	// bidirectional channels to unidirectional channels are allowed.
	// the other way around is not allowed.
	go counter(naturals)
	go squarer(naturals, squares)

	printer(squares)
}
