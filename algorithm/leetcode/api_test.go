package leetcode

//二叉树 中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BTreeToArray TreeNode->[]int
func (*TreeNode) BTreeToArray(tree *TreeNode) []int {
	return inorderTraversal(tree)
}

//SortArrayToBTree  []int->TreeNode
func (*TreeNode) SortedArrayToBTree(array []int) *TreeNode {
	return sortedArrayToBST(array)
}

//双向链表
type DoublyLink struct {
	Prev  *DoublyLink
	Value int
	Key   int
	Next  *DoublyLink
}

//移除掉Node
func (*DoublyLink) Remove(node *DoublyLink, setNull bool) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if setNull {
		node = nil
	}
}
