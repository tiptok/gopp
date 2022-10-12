package leetcode

import (
	"fmt"
	"log"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type In struct {
		A *ListNode
	}
	input := []In{
		In{A: NewListNode(2).Add(4, 3)},
		In{A: NewListNode(5).Add(1, 3, 6)},
		In{A: NewListNode(3).Add(4, 1)},
	}

	for _, in := range input {
		log.Println("---------------New Test-------------")
		log.Printf("Input: ( %s ) \n", in.A.PrintResult())
		ouput := reverseList(in.A)
		log.Printf("Output: %s\n", ouput.PrintResult())
	}
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func Test_0002(t *testing.T) {
	log.Println(fmt.Sprintf("%.2f", float64(101)/float64(100.0)))
}
