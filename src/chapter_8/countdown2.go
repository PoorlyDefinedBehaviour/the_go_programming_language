package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	// pretend to do something
}

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // wait for any key press
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10 * time.Second):
		launch()
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}
