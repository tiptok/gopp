package leetcode

//二叉树 中序遍历

// 双向链表
type DoublyLink struct {
	Prev  *DoublyLink
	Value int
	Key   int
	Next  *DoublyLink
}

// 移除掉Node
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
