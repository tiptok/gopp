package leetcode

import (
	"github.com/tiptok/gopp/algorithm/leetcode/structures"
	"log"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	input := [][]int{
		{-10, 3, 0, 5, 9},
		{1, 2, 3, 4, 5},
	}
	for i := range input {
		out := structures.SortedArrayToBST(input[i])
		log.Println("Input:", input[i], " output:", &out)
	}
}
