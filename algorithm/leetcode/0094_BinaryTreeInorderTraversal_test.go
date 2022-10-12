package leetcode

import (
	"log"
	"testing"
)

//二叉树 中序遍历
func Test_inorderTraversal(t *testing.T) {
	input := [][]int{
		{-10, 3, 0, 5, 9},
		{1, 2, 3, 4, 5},
	}
	for i := range input {
		inBST := SortedArrayToBST(input[i])
		out := postorderTraversal(inBST)
		log.Println("Input:", input[i], " output:", out)
	}
}

//使用堆栈的方式
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	var (
		cur *TreeNode = root
	)

	for cur != nil || (len(stack) != 0 && stack != nil) {
		if cur != nil {
			stack = Push(stack, cur)
			cur = cur.Left
		} else {
			stack, cur = Pop(stack)
			ret = append(ret, cur.Val)
			cur = cur.Right
		}
	}
	return ret
}

func Pop(stack []*TreeNode) ([]*TreeNode, *TreeNode) {
	if len(stack) == 0 || stack == nil {
		return stack, nil
	}
	if len(stack) == 1 {
		return nil, stack[0]
	}
	p := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	log.Println("Pop->", p.Val)
	return stack, p
}
func Push(stack []*TreeNode, cur *TreeNode) []*TreeNode {
	if stack == nil {
		stack = make([]*TreeNode, 0)
	}
	log.Println("Push->", cur.Val)
	stack = append(stack, cur)
	return stack
}

//二叉树 前序遍历
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	var (
		cur *TreeNode = root
	)
	for cur != nil || (len(stack) != 0 && stack != nil) {
		if cur != nil {
			ret = append(ret, cur.Val)
			stack = Push(stack, cur)
			cur = cur.Left
		} else {
			stack, cur = Pop(stack)
			if cur != nil {
				cur = cur.Right
			}
		}
	}
	return ret
}

//二叉树 后序遍历   遍历:左右中
/*
		 0
	-10     5
	   3      9
*/
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	var (
		cur *TreeNode = root
	)
	stack = Push(stack, cur)
	for len(stack) != 0 && stack != nil {
		var node *TreeNode
		stack, node = Pop(stack)
		if node.Left != nil {
			stack = Push(stack, node.Left)
		}
		if node.Right != nil {
			stack = Push(stack, node.Right)
		}
		ret = append(ret, node.Val)
	}

	//方向输出
	l := len(ret)
	for i := 0; i < l/2; i++ {
		tmp := ret[i]
		ret[i] = ret[l-i-1]
		ret[l-i-1] = tmp
	}
	return ret
}

//预排序遍历树 https://my.oschina.net/XYleung/blog/99604
