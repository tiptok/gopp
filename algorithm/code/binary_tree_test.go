package code

// 二叉树
// 104.maximum-depth-of-binary-tree]
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// 返回最长的所有元素
func maxDepth2(root *TreeNode) int {
	res := maxDepth2Helper(root)
	return len(res)
}

func maxDepth2Helper(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	left := maxDepth2Helper(root.Left)
	right := maxDepth2Helper(root.Right)

	if len(left) > len(right) {
		res = append(res, left...)
		return res
	}
	res = append(res, right...)
	return res
}

//type ListNode struct {
//     Val int
//     Next *ListNode
//}
func deleteDuplicates(head *ListNode) *ListNode {
	var fakeHead *ListNode = &ListNode{Val: 0}
	fakeHead.Next = head
	pre := fakeHead
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return fakeHead.Next
}
