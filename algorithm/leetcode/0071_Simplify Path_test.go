package leetcode

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("Cannot pop from an empty stack")
	}
	value := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return value
}

// 解法一
func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	var stack Stack[string]
	for _, cur := range list {
		switch cur {
		case "", ".":
			break
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		default:
			stack.Push(cur)

		}
	}
	if stack.IsEmpty() {
		return "/"
	}

	var res strings.Builder
	for _, item := range stack {
		res.WriteString("/")
		res.WriteString(item)
	}
	return res.String()
}

// 解法三
func simplifyPath3(path string) string {
	list := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(list); i++ {
		cur := list[i]
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}

// 解法二 golang 的官方库 API
func simplifyPath1(path string) string {
	return filepath.Clean(path)
}

func Test_simplifyPath(t *testing.T) {
	assert.Equal(t, "/home/user/Pictures", simplifyPath("/home/user/Documents/../Pictures"))
}
