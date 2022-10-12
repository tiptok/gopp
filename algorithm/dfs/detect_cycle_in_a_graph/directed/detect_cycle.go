package main

import "fmt"

//## https://yuminlee2.medium.com/detect-cycle-in-a-graph-4461b6000845

func main() {
	edges := [][]int{
		{0, 1},
		{0, 3},
		{0, 5},
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
		{5, 3},
	}
	graph, nodes := buildGraph(edges)
	fmt.Println("detect cycle result:", detect(nodes, graph, fmt.Printf))
}
