package leetcode

import (
	"log"
	"testing"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func TestRemoveNthFromEnd(t *testing.T) {
	type In struct {
		A *ListNode
		N int
	}

	input := []In{
		In{A: NewListNode(2).Add(4, 3, 7, 8), N: 2},
		In{A: NewListNode(2).Add(4, 3, 6), N: 2},
		//In{A:NewListNode(2),N:1},
		In{A: NewListNode(2).Add(2), N: 2},
		In{A: NewListNode(2).Add(2), N: 1},
		In{A: nil, N: 2},
	}

	for _, in := range input {
		log.Println("---------------New Test-------------")
		log.Printf("Input: ( %s ) N:%d \n", in.A.PrintResult(), in.N)
		ouput := removeNthFromEnd(in.A, in.N)
		log.Printf("Output: %s\n", ouput.PrintResult())
	}
}

//type ListNode struct {
//	     Val int
//	     Next *ListNode
//}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	var dumpHead *ListNode = &ListNode{Val: 0}
	dumpHead.Next = head
	/*双指针 偏移  p q 相距 n+1*/
	p := dumpHead
	q := dumpHead
	above := n + 1
	for {
		if above <= 0 {
			p = p.Next
		}
		if q.Next == nil {
			delNode := p.Next
			p.Next = delNode.Next
			break
		}
		q = q.Next
		above--
	}
	return dumpHead.Next
}
