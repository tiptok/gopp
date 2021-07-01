package code

import (
	"reflect"
	"testing"
)

// 344. Reverse String
func Test_reverseString(t *testing.T) {
	input := []byte("hello")
	excecpt := []byte("olleh")
	reverseString2(input)
	if !reflect.DeepEqual(input, excecpt) {
		t.Fatalf("out : %v except:%v", input, excecpt)
	}
}

func reverseString(s []byte) {
	var out *[]byte = new([]byte)
	reverse(s, 0, out)
	for i := 0; i < len(s); i++ {
		s[i] = (*out)[i]
	}
}

func reverse(s []byte, index int, out *[]byte) {
	if index == len(s) {
		return
	}
	reverse(s, index+1, out)
	*out = append(*out, s[index])
}

func reverseString2(s []byte) {
	i, j := 0, len(s)-1
	var tmp byte
	for {
		if !(i < j) {
			break
		}
		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
}

// 24. Swap Nodes in Pairs
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next
	// 翻转当前阶段指针
	next := head.Next
	next.Next = head
	head.Next = swapPairs(nextHead)
	return next
}

// 509.Fibonacci Number
func Test_fib(t *testing.T) {
	input := 2
	excecpt := 1
	out := fib(input)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}

var dp = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		return n
	}
	if dp[n] != 0 {
		return dp[n]
	}
	res := fib(n-2) + fib(n-1)
	dp[n] = res
	return res
}
