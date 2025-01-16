package structures

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BTreeToArray TreeNode->[]int
func (*TreeNode) BTreeToArray(tree *TreeNode) []int {
	return inorderTraversal(tree)
}

// SortArrayToBTree  []int->TreeNode
func (*TreeNode) SortedArrayToBTree(array []int) *TreeNode {
	return sortedArrayToBST(array)
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

// 二分法
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

// 使用堆栈的方式
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

func PreorderTraversal(root *TreeNode) []int {
	return preorderTraversal(root)
}

// 二叉树 前序遍历
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
