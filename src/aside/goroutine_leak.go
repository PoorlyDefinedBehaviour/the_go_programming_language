package main

import (
	"fmt"
	"time"
)

// from https://dave.cheney.net/2016/12/22/never-start-a-goroutine-without-knowing-how-it-will-stop

func f(channel <-chan int) {
	defer fmt.Println("leaving f()")

	go func() {
		defer fmt.Println("leaving goroutine")
		// This goroutine does not stop until the channel is closed.
		// We actually wanted this goroutine to run until f() is done.
		for value := range channel {
			fmt.Println(value)
		}
	}()

	for i := 0; i < 1000; i++ {
		// do some work
	}
}

func main() {
	channel := make(chan int)

	f(channel)

	time.Sleep(1 * time.Second)

	fmt.Println("closing channel")
	close(channel)

	time.Sleep(1 * time.Second)
}
