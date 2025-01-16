package leetcode

import (
	"github.com/tiptok/gopp/algorithm/leetcode/structures"
	"log"
	"testing"
)

// 二叉树 中序遍历
func Test_inorderTraversal(t *testing.T) {
	input := [][]int{
		{-10, 3, 0, 5, 9},
		{1, 2, 3, 4, 5},
	}
	for i := range input {
		inBST := structures.SortedArrayToBST(input[i])
		out := structures.PreorderTraversal(inBST)
		log.Println("Input:", input[i], " output:", out)
	}
}
