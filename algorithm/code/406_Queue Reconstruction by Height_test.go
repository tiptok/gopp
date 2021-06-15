package main

import (
	"sort"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	var input = [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	result := reconstructQueue(input)
	t.Log(result)
}

func reconstructQueue(people [][]int) [][]int {
	var result [][]int
	var height []int
	var mapPeople = make(map[int][]int)
	if len(people) == 0 {
		return result
	}

	for i := 0; i < len(people); i++ {
		p := people[i]
		if _, ok := mapPeople[p[0]]; ok {
			mapPeople[p[0]] = append(mapPeople[p[0]], p[1])
		} else {
			mapPeople[p[0]] = []int{p[1]}
			height = append(height, p[0])
		}
	}

	sort.Ints(height)

	for i := len(height) - 1; i >= 0; i-- {
		h := height[i]
		sort.Ints(mapPeople[h])
		for j := 0; j <= len(mapPeople[h])-1; j++ {
			p := mapPeople[h][j]
			if len(result) == 0 {
				result = [][]int{{h, p}}
				continue
			}
			rest := append([][]int{{h, p}}, result[p:]...)
			result = append(result[:p], rest...)
		}
	}
	return result
}
