package main

func detect(nodes []int, graph map[int][]int, output func(format string, values ...interface{}) (n int, err error)) bool {
	visiting := map[int]bool{}
	visited := map[int]bool{}

	for _, node := range nodes {
		output("------------------------------------------\n")
		output("dfs node %v \n", node)
		if detectCycle(graph, node, visiting, visited, output) {
			return true
		}
	}
	return false
}

func detectCycle(graph map[int][]int,
	node int,
	visiting, visited map[int]bool, output func(format string, values ...interface{}) (n int, err error)) bool {
	if _, found := visited[node]; found {
		output("node %d is already visited(black) -> skip\n", node)
		return false
	}

	if _, found := visiting[node]; found {
		output("node %d is in visiting(gray) -> a cycle is detected\n\n", node)
		return true
	}
	visiting[node] = true
	output("nodes in visiting(gray): %+v\n", visiting)
	output("nodes in visited(black): %+v\n\n", visited)

	for _, descendant := range graph[node] {
		output("current node: node %d\n", node)
		output("visit descendant: node %d\n", descendant)
		if detectCycle(graph, descendant, visiting, visited, output) {
			return true
		}
	}

	delete(visiting, node)
	visited[node] = true
	output("finish explore node %d\n", node)
	output("nodes in visiting(gray): %+v\n", visiting)
	output("nodes in visited(black): %+v\n\n", visited)
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
		graph[a] = append(graph[a], b)
	}
	return graph, nodes
}
