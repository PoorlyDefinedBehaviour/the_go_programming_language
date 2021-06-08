package main

import (
	"fmt"
	"time"
)

func launch() {
	// pretend to do something
}

func main() {
	fmt.Println("Commencing countdown.")

	tick := time.Tick(1 * time.Second)

	for coundtdown := 10; coundtdown > 0; coundtdown-- {
		fmt.Println(coundtdown)
		<-tick
	}

	launch()
}
