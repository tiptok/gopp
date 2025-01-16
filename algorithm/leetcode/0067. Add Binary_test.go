package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

/*
Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
题目大意 #
给你两个二进制字符串，返回它们的和（用二进制表示）。输入为 非空 字符串且只包含数字 1 和 0
*/
func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	if aLen < bLen {
		a, b, aLen, bLen = b, a, bLen, aLen
	}
	var (
		res   = ""
		carry = 0
	)
	for i := 0; i < aLen; i++ {
		var sum = 0
		sum += int(a[aLen-1-i] - '0')
		if i < bLen {
			sum += int(b[bLen-1-i] - '0')
		}
		sum += carry
		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		res = strconv.Itoa(sum%2) + res
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}

func TestAddBinary(t *testing.T) {
	//assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "100", addBinary("11", "1"))
}
