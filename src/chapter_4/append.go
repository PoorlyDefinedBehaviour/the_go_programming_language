package main

import "fmt"

// O(1) append amortized append time
func appendInt(xs []int, x int) []int {
	var newxs []int
	newxslen := len(xs) + 1

	if newxslen <= cap(xs) {
		newxs = xs[:newxslen]
	} else {
		newxscap := newxslen
		if newxscap < 2*len(xs) {
			newxscap = 2 * len(xs)
		}

		newxs = make([]int, newxslen, newxscap)
		copy(newxs, xs)
	}

	newxs[len(xs)] = x

	return newxs
}

func main() {
	xs := []int{}
	ys := []int{}

	for i := 0; i < 10; i += 1 {
		ys = appendInt(xs, i)

		fmt.Printf("%d	cap=%d\t%v\n", i, cap(ys), ys)

		xs = ys
	}
}
