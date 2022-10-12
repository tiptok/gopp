package leetcode

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

/*slice*/
func TestAddTwoNum(t *testing.T) {
	type In struct {
		A []int
		B []int
	}

	input := []In{
		In{A: []int{2, 4, 3}, B: []int{5, 6, 4}},
		In{A: []int{2, 4, 3, 6}, B: []int{5, 6, 4}},
		In{A: []int{2, 4, 3}, B: []int{5, 6, 4, 3}},
	}

	Print := func(in []int) string {
		outBuffer := bytes.NewBuffer(nil)
		for i, inVal := range in {
			outBuffer.WriteString(fmt.Sprintf("%d", inVal))
			if i < (len(in) - 1) {
				outBuffer.WriteString(fmt.Sprintf("%s", " -> "))
			}
		}
		return outBuffer.String()
	}

	for _, in := range input {
		ouput := AddTwoNum(in.A, in.B)
		log.Println("---------------New Test-------------")
		log.Printf("Input: ( %s ) + ( %s ) \n", Print(in.A), Print(in.B))
		log.Printf("Output: %s\n", Print(ouput))
	}
}

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
func AddTwoNum(listA, listB []int) []int {
	lenA := len(listA)
	lenB := len(listB)
	var (
		total int = 0
	)
	if lenA > lenB {
		total = lenA
	} else {
		total = lenB
	}
	listRet := make([]int, total)
	carried := 0
	for i := 0; i < total; i++ {
		valA, valB := 0, 0
		if i < lenA {
			valA = listA[i]
		}
		if i < lenB {
			valB = listB[i]
		}
		valSum := valA + valB
		listRet[i] = valSum%10 + carried
		carried = valSum / 10
	}
	return listRet
}

/*list*/
func TestAddTwoNumListNode(t *testing.T) {
	type In struct {
		A *ListNode
		B *ListNode
	}

	input := []In{
		In{A: NewListNode(2).Add(4, 3), B: NewListNode(5).Add(6, 4)},
		In{A: NewListNode(2).Add(4, 3, 6), B: NewListNode(5).Add(6, 4)},
		In{A: NewListNode(2).Add(4, 3), B: NewListNode(5).Add(6, 4, 3)},
		In{A: NewListNode(5), B: NewListNode(5)},
		In{A: NewListNode(1), B: NewListNode(9).Add(9)},
	}

	for _, in := range input {
		ouput := addTwoNumbers(in.A, in.B)
		log.Println("---------------New Test-------------")
		log.Printf("Input: ( %s ) + ( %s ) \n", in.A.PrintResult(), in.B.PrintResult())
		log.Printf("Output: %s\n", ouput.PrintResult())
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}
func NewListNodes(val ...int) *ListNode {
	var listNode *ListNode = NewListNode(0)
	listNode.Add(val...)
	return listNode.Next
}
func (l *ListNode) Add(val ...int) *ListNode {
	next := l
	for {
		if next.Next == nil {
			for _, v := range val {
				next.Next = &ListNode{
					Val: v,
				}
				next = next.Next
			}
			break
		}
		next = next.Next
	}
	return l
}
func (l *ListNode) PrintResult() string {
	outBuf := bytes.NewBuffer(nil)
	for node := l; node != nil; node = node.Next {
		outBuf.WriteString(fmt.Sprintf("%d", node.Val))
		if node.Next != nil {
			outBuf.WriteString(fmt.Sprintf("%s", " -> "))
		}
	}
	return outBuf.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmpNode *ListNode = &ListNode{}
	rspNode := tmpNode
	carried := 0
	for {
		valSum := l1.Val + l2.Val + carried
		tmpNode.Val = valSum % 10
		carried = valSum / 10
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil && carried == 0 {
			break
		}
		if l1 == nil {
			l1 = &ListNode{}
		}
		if l2 == nil {
			l2 = &ListNode{}
		}
		tmpNode.Next = &ListNode{}
		tmpNode = tmpNode.Next
	}
	return rspNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var cN1, cN2 *ListNode = l1, l2
	var tmpNode *ListNode = &ListNode{}
	rspNode := tmpNode
	carried := 0
	for {
		valSum := cN1.Val + cN2.Val + carried
		tmpNode.Val = valSum % 10
		carried = valSum / 10
		//log.Println("tmpNode 地址:",tmpNode)
		cN1 = cN1.Next
		cN2 = cN2.Next
		if cN1 == nil && cN2 == nil && carried == 0 {
			break
		}
		if cN1 == nil {
			cN1 = &ListNode{}
		}
		if cN2 == nil {
			cN2 = &ListNode{}
		}
		tmpNode.Next = &ListNode{}
		tmpNode = tmpNode.Next
	}
	//log.Println("-------指针地址----:")
	//log.Println("地址1:",rspNode)
	//log.Println("地址2:",tmpNode)
	return rspNode
}
