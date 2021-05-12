package main

import "fmt"

func main() {
	ages1 := make(map[string]int) // mapping from strings to int

	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	emptyMap := map[string]int{}

	fmt.Println(ages1)
	fmt.Println(ages2)
	fmt.Println(emptyMap)

	fmt.Println(ages2["alice"])

	delete(ages2, "alice")

	fmt.Println(ages2)
}
