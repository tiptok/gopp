package leetcode

import (
	"log"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	input := [][]int{
		{-10, 3, 0, 5, 9},
		{1, 2, 3, 4, 5},
	}
	for i := range input {
		out := sortedArrayToBST(input[i])
		log.Println("Input:", input[i], " output:", &out)
	}
}
func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return CreateBTree(nums, 0, len(nums)-1)
}

//二分法
func CreateBTree(nums []int, lIndex, rIndex int) *TreeNode {
	if lIndex <= rIndex {
		mid := (lIndex + rIndex) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = CreateBTree(nums, lIndex, mid-1)
		node.Right = CreateBTree(nums, mid+1, rIndex)
		return node
	} else {
		return nil
	}
}
