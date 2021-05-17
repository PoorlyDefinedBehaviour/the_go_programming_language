package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
Many programming language implementations use a fixed-size function call stack.
Sizes from 64KB to 2MB are typical. Fixed-size stacks impose a limit on the depth
of recursion. In constrat, typical Go implementations use variable-size stacks
that start small and grow as needed up to a limit on the order of a gigabyte.
*/
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	outline(nil, doc)
}

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		outline(stack, child)
	}
}
