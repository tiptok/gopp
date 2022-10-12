package main

import (
	"fmt"
)

//## https://yuminlee2.medium.com/detect-cycle-in-a-graph-4461b6000845

func main() {
	inputs := []struct {
		edges  [][]int
		circle bool
		name   string
	}{
		{
			edges: [][]int{
				{0, 1},
				{0, 3},
				{0, 5},
				{1, 2},
				{1, 3},
				{3, 4},
				{4, 5},
				{5, 3},
			},
			circle: true,
			name:   "用例有环",
		},
		{
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 5},
				{2, 3},
				{2, 4},
			},
			circle: false,
			name:   "用例无环",
		},
	}

	for _, input := range inputs {
		graph, nodes := buildGraph(input.edges)
		fmt.Println("graphL", graph, "node:", nodes)
		got := detect(nodes, graph, fmt.Printf)
		fmt.Println("detect cycle result:", got)
		if got != input.circle {
			fmt.Printf("title:%v except %v got %v", input.name, input.circle, got)
		}
	}

}
