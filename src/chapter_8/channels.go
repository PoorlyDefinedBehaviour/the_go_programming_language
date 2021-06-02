package main

/*
A channel is a communication mechanism that lets one goroutine
send values to another goroutine. Each channel is a conduit for values
of a particular type, called the channel's element type.
The type of a channel whose elements have type int is written:
chan int.

To create a channel, we use the built-in make function:

channel := make(chan int) // channel has type chan int

channel is a reference to the data structure created by make.
When we copy a channel or pass one as argument to a function,
we are copying a reference.

As with other reference types, the zero value of a channel is nil.

Channels of the same type may be compared using ==. The comparison
is true if both are references to the same channel data structure.

A channel may also be compared to nil.

A channel has two principal operations: send an receive.

A send statemen transmit a value from one goroutine, though the cannel,
to another goroutine executing a corresponding receive expression. Both operations
are written using the <- operator.

channel <- x  // send statement
x = <-channel // a receive expression in an assignment statement
<-channel     // a receive statement, result is unused.

Channels support a third operation, close, which sets a flag indicating
that no more values will ever be sent on this channel.
Subsequent attemps to send will panic.
Receive operations on a closed channel yield the values that have been sent
until no more values are left. Any receive operations thereafter complete
immediately and yield the zero value of the channel's element type.

To close a channel, we call the built-in close function:

close(channel)

A channel created with a simple call to make is called
an unbuffered channel, but make accepts an optional second argument,
an integer called the channel's capacity. If the capcity is greater than zero,
makes creates a buffer channel.

channel := make(chan int)    // unbuffered channel
channel := make(chan int, 0) // unbuffer channel
channel := make(chan int, 3) // buffered channel with capacity 3

Unbuffered Channels

A send operation on an unbuffered channel blocks the sending goroutine
until another goroutine executes a corresponding receive on the same channel.

If the receive operation was attempted first, the receiving goroutine is blocked
until another goroutine performs a send on the same channel.

Basically, unbuffered channels block whoever uses them.

When a value is sent on an unbuffered channel, the value is received
by the other goroutine and then the sender unblocks.
*/

func main() {

}
