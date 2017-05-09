package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdges(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdges(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdges("a", "b")
	addEdges("c", "d")
	addEdges("c", "b")
	fmt.Println(hasEdges("a", "e"))
	fmt.Println(hasEdges("a", "c"))
	fmt.Println(hasEdges("a", "b"))
	fmt.Println(hasEdges("f", "e"))
	fmt.Println(hasEdges("c", "b"))
	fmt.Println(hasEdges("c", "d"))
	fmt.Println(hasEdges("b", "a"))

}
