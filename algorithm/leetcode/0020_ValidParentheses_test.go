package leetcode

import (
	"fmt"
	"log"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	input := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"((",
		"(",
	}
	for i := 0; i < len(input); i++ {
		out := isValid(input[i])
		log.Println(fmt.Sprintf("Input:%v Output:%v", input[i], out))
	}
}

//使用栈
func isValid(s string) bool {
	buf := make([]byte, 0)
	valid := []string{"[]", "{}", "()"}

	checkValid := func(a, b byte) bool {
		for i := 0; i < len(valid); i++ {
			if valid[i] == fmt.Sprintf("%v%v", string(a), string(b)) {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(s); i++ {
		j := len(buf)
		if j >= 1 {
			if checkValid(buf[j-1], s[i]) {
				buf = buf[:(j - 1)]
			} else {
				buf = append(buf, s[i])
			}
		} else {
			buf = append(buf, s[i])
		}
	}
	if len(buf) == 0 {
		return true
	} else {
		return false
	}
}
