package main

import (
	"fmt"
	"time"
)

func counter(next chan int) {
	for i := 0; i < 5; i++ {
		next <- i
	}
	close(next)
}

func square(previous, next chan int) {
	for x := range previous {
		next <- x * x
	}
	close(next)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go square(naturals, squares)

	for x := range squares {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
