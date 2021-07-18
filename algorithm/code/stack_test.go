package code

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 155 Min Stack
// 使用链表
type MinStack struct {
	head *MSNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	minVal := s.GetMin()
	if s.head == nil {
		s.head = NewMSNode(val, val, nil)
	} else {
		s.head = NewMSNode(val, min(minVal, val), s.head)
	}
}

func (s *MinStack) Pop() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
}

func (s *MinStack) Top() int {
	if s.head == nil {
		return 0
	}
	return s.head.val
}

func (s *MinStack) GetMin() int {
	if s.head == nil {
		return 1 << 31
	}
	return s.head.min
}

type MSNode struct {
	val  int
	min  int
	next *MSNode
}

func NewMSNode(val, min int, next *MSNode) *MSNode {
	return &MSNode{
		val:  val,
		min:  min,
		next: next,
	}
}

// two stack
// 1.改造为单个stack  item{min int,val int}
// type MinStack struct {
// 	min   []int
// 	stack []int
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{
// 		min:   make([]int, 0),
// 		stack: make([]int, 0),
// 	}
// }

// func (s *MinStack) Push(val int) {
// 	min := s.GetMin()
// 	if min < val {
// 		s.min = append(s.min, min)
// 	} else {
// 		s.min = append(s.min, val)
// 	}
// 	s.stack = append(s.stack, val)
// }

// func (s *MinStack) Pop() {
// 	if len(s.stack) == 0 {
// 		return
// 	}
// 	s.stack = s.stack[:len(s.stack)-1]
// 	s.min = s.min[:len(s.min)-1]
// }

// func (s *MinStack) Top() int {
// 	if len(s.stack) == 0 {
// 		return 0
// 	}
// 	return s.stack[len(s.stack)-1]
// }

// func (s *MinStack) GetMin() int {
// 	if len(s.min) == 0 {
// 		return 1 << 31
// 	}
// 	return s.min[len(s.min)-1]
// }

func Test_MinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	//assert.Equal(t, minStack.stack, []int{-2, 0, -3})
	assert.Equal(t, minStack.GetMin(), -3)
	minStack.Pop()
	assert.Equal(t, minStack.Top(), 0)
	assert.Equal(t, minStack.GetMin(), -2)
}

// 150.Evaluate Reverse Polish Notation  逆波兰表达式求值
func Test_evalRPN(t *testing.T) {
	input := []string{"2", "1", "+", "3", "*"}
	excecpt := 9
	out := evalRPN(input)
	assert.Equal(t, excecpt, out)
}
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			length := len(stack)
			if length < 2 {
				return -1
			}
			a := stack[length-2]
			b := stack[length-1]
			stack = stack[:length-2]
			var result int
			switch tokens[i] {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			stack = append(stack, result)
		default:
			vals, _ := strconv.Atoi(tokens[i])
			stack = append(stack, vals)
		}
	}
	return stack[0]
}
