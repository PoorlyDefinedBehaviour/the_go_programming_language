package main

import "fmt"

// equivalent to switch true { ... }
func signum(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func main() {
	fmt.Println(signum(-1))
	fmt.Println(signum(0))
	fmt.Println(signum(1))
}
