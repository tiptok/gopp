package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string “”.

Input: strs = ["flower","flow","flight"]
Output: "fl"
*/
func longestCommonPrefix(strs []string) string {
	p := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(p) && i < len(s) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func longestCommonPrefix1(strs []string) string {
	var index = 0
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for {
		var redo = false
		for i := 0; i < len(strs)-1; i++ {
			if len(strs[i]) <= index {
				break
			}
			if len(strs[i+1]) <= index {
				break
			}
			if strs[i][index] != strs[i+1][index] {
				break
			}
			if i == len(strs)-2 {
				redo = true
			}
		}
		if !redo {
			break
		}
		index++
	}
	return strs[0][0:index]
}

func Test_longestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
