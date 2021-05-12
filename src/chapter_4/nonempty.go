package main

import "fmt"

func nonempty(strings []string) []string {
	result := strings[:0]

	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

func main() {
	xs := []string{"hello", "", "world", "", "!"}

	fmt.Println(nonempty(xs))
	fmt.Println(xs)
}
