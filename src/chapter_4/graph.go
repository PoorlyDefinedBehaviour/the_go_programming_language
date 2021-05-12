package main

import "fmt"

type Graph map[string]map[string]bool

func new() Graph {
	return make(Graph)
}

func (graph *Graph) addEdge(from string, to string) {
	edges := (*graph)[from]

	if edges == nil {
		edges = make(map[string]bool)
		(*graph)[from] = edges
	}

	edges[to] = true
}

func (graph *Graph) hasEdge(from string, to string) bool {
	return (*graph)[from][to]
}

func main() {
	graph := new()

	fmt.Println(graph)

	graph.addEdge("a", "b")
	graph.addEdge("b", "c")

	fmt.Println(graph)
	fmt.Println(graph.hasEdge("a", "b"))
	fmt.Println(graph.hasEdge("b", "a"))
	fmt.Println(graph.hasEdge("b", "c"))
}
