package leetcode

import (
	"fmt"
	"log"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type In struct {
		A *ListNode
		B *ListNode
	}
	Input := []*In{
		{A: NewListNode(0).Add(3, 4, 7, 8), B: NewListNode(1).Add(2, 2, 3)},
		{A: NewListNode(0).Add(1, 2, 2, 6), B: NewListNode(2).Add(3, 4, 7, 8)},
	}
	for i, _ := range Input {
		out := mergeTwoLists2(Input[i].A, Input[i].B)
		PrintResult(out.PrintResult(), "", Input[i].A.PrintResult(), Input[i].B.PrintResult())
	}
}

func PrintResult(out interface{}, want interface{}, in ...interface{}) {
	log.Println(fmt.Sprintf("In:%v Out:%v Want:%v", in, out, want))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var retList *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		retList = l1
		retList.Next = mergeTwoLists(l1.Next, l2)
	} else {
		retList = l2
		retList.Next = mergeTwoLists(l1, l2.Next)
	}
	return retList
}

//使用递归算法 合并两个已排序链表 ---->>faster
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}
