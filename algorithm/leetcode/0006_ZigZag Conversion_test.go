package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下
P   A   H   N
A P L S I I G
Y   I   R

|    /|    /|    /|
|  /  |  /  |  /  |
|/    |/    |/    |
*/
func convert(s string, numRows int) string {
	result := make([][]byte, numRows)
	down, up := 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			result[down] = append(result[down], s[i])
			down += 1
			i++
		} else if up > 0 {
			result[up] = append(result[up], s[i])
			up -= 1
			i++
		} else {
			down, up = 0, numRows-2
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range result {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}

func Test_ZigZag(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
}
