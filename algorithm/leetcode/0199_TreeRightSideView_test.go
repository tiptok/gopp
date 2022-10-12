package leetcode

import (
	"log"
	"testing"
)

func Test_tmpqueue(t *testing.T) {

	input := [][]int{
		{1, 2, 3, 4, 5},
		//{-10,3,0,5,9},
	}
	for i := range input {
		tree := new(TreeNode)
		out := rightSideView(tree.SortedArrayToBTree(input[i]))
		log.Println("in:", tree, "out:", out)
	}
}

//二叉树右视图
func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	//result = append(result,root.Val)
	queue = append(queue, root)
	//分层计算最右节点
	for len(queue) > 0 {
		tmpqueue := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tmpqueue = append(tmpqueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpqueue = append(tmpqueue, queue[i].Right)
			}
		}
		result = append(result, queue[len(queue)-1].Val)
		queue = queue[0:0]
		queue = append(queue, tmpqueue...)
	}
	return result
}
