package structures

import (
	"bytes"
	"fmt"
)

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
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
