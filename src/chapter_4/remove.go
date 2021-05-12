package main

import "fmt"

func remove(slice []int, i int) []int {
	if len(slice) <= 1 {
		return []int{}
	}

	result := make([]int, len(slice))

	copy(result, slice)
	result[i] = result[len(result)-1]
	return result[:len(result)-1]
}

func main() {
	xs := []int{1, 2, 3, 4, 5}

	fmt.Println(xs)
	fmt.Println(remove(xs, 2))
	fmt.Println(xs)
}
