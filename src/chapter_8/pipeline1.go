package main

import (
	"fmt"
	"time"
)

func counter(next chan int) {
	for i := 0; ; i++ {
		next <- i
	}
}

func square(previous, next chan int) {
	for {
		x := <-previous
		next <- x * x
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go square(naturals, squares)

	for {
		fmt.Println(<-squares)
		time.Sleep(1 * time.Second)
	}
}
