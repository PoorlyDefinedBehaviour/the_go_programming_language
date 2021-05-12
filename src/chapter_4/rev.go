package main

import "fmt"

func reverse(s []int) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}

func main() {
	// when a slice literal is used, an array variable of the right size
	// is created implicitly and a slice that points to it is yielded.
	xs := []int{0, 1, 2, 3, 4, 5}

	reverse(xs[:])

	// xs is modified even though a slice is passed to reverse(...)
	// because a slice is just a pointer, a length and a capacity.
	fmt.Println(xs) // [5 4 3 2 1 0]
}
