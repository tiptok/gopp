package leetcode

import (
	"github.com/stretchr/testify/assert"
	"github.com/tiptok/gopp/algorithm/leetcode/structures"
	"testing"
)

/*
给定 2 个链表中结点的位置 m, n，反转这个两个位置区间内的所有结点。

Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

Input: head = [5], left = 1, right = 1
Output: [5]

解题思路 用一个cur指向当前，把当前的下一个节点移动到 pre.next，重复n-m次
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || head.Next == nil || left > right {
		return head
	}
	var newHead *ListNode = &ListNode{
		Val:  0,
		Next: head,
	}
	pre := newHead
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

/*
1 2 3 4 5
1 3 2 4 5   2 4 5
1 4 3 2 5   2 5
1 5 4 3 2   2
*/
func TestReverseBetween(t *testing.T) {
	assert.Equal(t, structures.List2Ints(reverseBetween(structures.Ints2List([]int{1, 2, 3, 4, 5}), 2, 5)), []int{1, 5, 4, 3, 2})
}
