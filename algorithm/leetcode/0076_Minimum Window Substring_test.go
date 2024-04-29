package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	if len(s) < len(t) {
		return ""
	}
	var tList, sList [256]int
	for i := 0; i < len(t); i++ {
		tList[val(t, i)]++
	}
	var res string
	var left, right, count, minW = 0, -1, 0, len(s) + 1
	var resLeft, resRight = -1, 0
	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sList[val(s, right+1)]++
			if sList[val(s, right+1)] <= tList[val(s, right+1)] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				resLeft = left
				resRight = right
			}
			if sList[val(s, left)] == tList[val(s, left)] {
				count--
			}
			sList[val(s, left)]--
			left++
		}
	}
	if resLeft != -1 {
		res = string(s[resLeft : resRight+1])
	}
	return res
}

func val(s string, idx int) byte {
	b := s[idx] - 'A'
	return b
}

func Test_minWindow(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
}
