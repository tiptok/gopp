package main

import "fmt"

type outputFunc func(format string, values ...interface{}) (n int, err error)

func detect(nodes []int, graph map[int][]int, output outputFunc) bool {
	visited := map[int]bool{}

	for _, node := range nodes {
		//node := nodes[0]
		output("------------------------------------------\n")
		output("dfs node %v \n", node)
		if detectCycle(graph, node, visited, -1, output) {
			return true
		}
		if len(visited) == len(nodes) {
			return false
		}
	}
	return false
}

func detectCycle(graph map[int][]int,
	node int, visited map[int]bool, parent int, output outputFunc) bool {
	if _, found := visited[node]; found {
		output("node %d has been visited\n\n", node)
		return true
	}

	visited[node] = true

	for _, descendant := range graph[node] {
		output("current node: %d\n", node)
		output("visited: %+v\n", visited)
		output("parent: %d\n", parent)
		if descendant != parent {
			fmt.Printf("visited descendant: node %d\n\n", descendant)
		} else {
			fmt.Printf("not visited descendant, node %d. It is parent.\n\n", descendant)
		}
		if descendant != parent && detectCycle(graph, descendant, visited, node, output) {
			return true
		}
	}

	return false
}

func buildGraph(edges [][]int) (map[int][]int, []int) {
	var graph = make(map[int][]int)
	var nodes []int
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if _, ok := graph[a]; !ok {
			nodes = append(nodes, a)
			graph[a] = make([]int, 0)
		}
		if _, ok := graph[b]; !ok {
			nodes = append(nodes, b)
			graph[b] = make([]int, 0)
		}
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph, nodes
}
