package main

import (
	"fmt"
	"sort"
)

// Go's sort package provides a function called Sort,
// that assumes nothing about the repesentation of either
// the sequence or its elements. Instead, it uses an interface,
// soprt.Interface, to specify the contract between the generic
// sort algorithm and each sequence type that may be sorted.
//
// Interfaces actually sound like a good ideia, but in reality,
// it is awful to use the sort package.
//
// How sort.Interface looks like:
//
// package sort
// type Interface interface {
// 	 Len() int
// 	 Less(i, j int) bool // i, j are indices
// 	 Swap(i, j int)
// }
//
// To sort a sequence, we need to define a type that implements
// these three methods, then apply sort.Sort to an instance
// of that type.

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"Charlie", "Bob", "Alice"}

	sort.Sort(StringSlice(names)) // could use sort.Strings(names) instead

	fmt.Println(names) // [Alice Bob Charlie]
}
