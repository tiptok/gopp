package leetcode

import (
	"fmt"
	"github.com/tiptok/gopp/algorithm/leetcode/structures"
	"testing"
)

type ListNode = structures.ListNode

/*
旋转链表 K 次。

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	total := 0
	cur := newHead
	for cur.Next != nil {
		cur = cur.Next
		total++
	}
	if (k % total) == 0 {
		return head
	}
	k = k % total
	// 首尾串联 1234512345
	cur.Next = head
	// 重新遍历
	cur = newHead
	for i := total - k; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	// 遍历结束截断当前的下个指针
	cur.Next = nil
	return res.Next
}

func TestRotateRight(t *testing.T) {
	qs := []question61{

		{
			para61{[]int{1, 2, 3, 4, 5}, 2},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{1, 2, 3, 4, 5}, 3},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{0, 1, 2}, 4},
			ans61{[]int{2, 0, 1}},
		},

		{
			para61{[]int{1, 1, 1, 2}, 3},
			ans61{[]int{1, 1, 2, 1}},
		},

		{
			para61{[]int{1}, 10},
			ans61{[]int{1}},
		},

		{
			para61{[]int{}, 100},
			ans61{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 61------------------------\n")

	for _, q := range qs {
		_, p := q.ans61, q.para61
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(rotateRight(structures.Ints2List(p.one), p.k)))
	}
	fmt.Printf("\n\n\n")
}

type question61 struct {
	para61
	ans61
}

// para 是参数
// one 代表第一个参数
type para61 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans61 struct {
	one []int
}

func Test_Problem61(t *testing.T) {

	qs := []question61{

		{
			para61{[]int{1, 2, 3, 4, 5}, 2},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{1, 2, 3, 4, 5}, 3},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{0, 1, 2}, 4},
			ans61{[]int{2, 0, 1}},
		},

		{
			para61{[]int{1, 1, 1, 2}, 3},
			ans61{[]int{1, 1, 2, 1}},
		},

		{
			para61{[]int{1}, 10},
			ans61{[]int{1}},
		},

		{
			para61{[]int{}, 100},
			ans61{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 61------------------------\n")

	for _, q := range qs {
		_, p := q.ans61, q.para61
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(rotateRight(structures.Ints2List(p.one), p.k)))
	}
	fmt.Printf("\n\n\n")
}
