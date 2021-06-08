package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A buffered channel has a queue of elements.
// The queue's maximum size is determined when it is created,
// by the capacity argument to make.
//
// make(chan int, capacity)
// make(chan int, 3) // channel has buffer of size 3.
//
// A send operation on a buffered channel inserts an element
// in the queue and a receive operation removes an element.
// If the channel is full, the send operation blocks its
// goroutine until space is made available by another
// goroutine's receive.
// If the channel is empty, the receive operation blocks its
// goroutine until a value is sent by another goroutine.
//
// Given a channel:
//
// channel := make(chan string, 3)
//
// We can send up to 3 values on this channel without the goroutine blocking:
//
// channel <- "A" // doesn't block
// channel <- "B" // doesn't block
// channel <- "C" // doesn't block
// channel <- "D" // blocks

func request(url string) string {
	rand.Seed(time.Now().UnixNano())

	// pretend to make a http request
	time.Sleep(time.Duration(rand.Intn(3)))

	return url
}

func getFastestResponse() string {
	responses := make(chan string, 3)

	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()

	return <-responses
}

func main() {
	fmt.Println(getFastestResponse())
}
