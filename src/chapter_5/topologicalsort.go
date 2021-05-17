package main

import (
	"fmt"
	"sort"
)

// The objective is to compute a sequence of courses
// that satisfies the prerequisite requirements of each one.
func main() {
	prerequisites := map[string][]string{
		"algorithms":            {"data structures"},
		"calculus":              {"linear algebra"},
		"compilers":             {"data structures", "formal languages", "compiler organization"},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range topologicalSort(prerequisites) {
		fmt.Printf("%d:\t%s\n", i, course)
	}
}

func mapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func topologicalSort(prerequisites map[string][]string) []string {
	order := make([]string, 0, len(prerequisites))
	seen := make(map[string]bool, len(prerequisites))

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(prerequisites[item])
				order = append(order, item)
			}
		}
	}

	keys := mapKeys(prerequisites)

	sort.Strings(keys)

	visitAll(keys)

	return order
}
