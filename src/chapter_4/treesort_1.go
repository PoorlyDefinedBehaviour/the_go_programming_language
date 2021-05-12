package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func newTree(value int) *tree {
	return &tree{value: value, left: nil, right: nil}
}

func (t *tree) add(value int) {
	currentNode := t

	for currentNode != nil {
		if value >= currentNode.value {
			if currentNode.right == nil {
				currentNode.right = newTree(value)
				return
			} else {
				currentNode = currentNode.right
			}
		} else {
			if currentNode.left == nil {
				currentNode.left = newTree(value)
				return
			} else {
				currentNode = currentNode.left
			}
		}
	}
}

func intoSlice(t *tree, slice *[]int) {
	if t == nil {
		return
	}

	intoSlice(t.left, slice)
	*slice = append(*slice, t.value)
	intoSlice(t.right, slice)
}

func (t *tree) toSlice() []int {
	slice := []int{}

	intoSlice(t, &slice)

	return slice
}

func main() {
	t := newTree(1)
	t.add(2)
	t.add(-2)

	fmt.Println(t.toSlice())
}
