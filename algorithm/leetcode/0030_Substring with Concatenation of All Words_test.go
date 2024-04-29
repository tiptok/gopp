package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]

给定一个源字符串 s，再给一个字符串数组，要求在源字符串中找到由字符串数组各种组合组成的连续串的起始下标，如果存在多个，在结果中都需要输出
*/

func TestFindSubstring(t *testing.T) {
	input1 := "barfoothefoobarman" //barfoothefoobarman
	input2 := []string{"bar", "foo"}
	out := findSubstring(input1, input2)
	assert.Equal(t, []int{0, 9}, out)
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}

	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	for i := 0; i < len(s)-totalLen+1; i++ {
		// key不存在提早退出
		if _, ok := counter[s[i:i+length]]; !ok {
			continue
		}
		tmpCounter = copyMap(counter)
		for j := i; j < i+totalLen && j+length-1 < len(s); j += length {
			key := s[j : j+length]
			if tmpCounter[key] > 0 {
				tmpCounter[s[j:j+length]]--
				if checkWords(tmpCounter) && (i+totalLen == j+length) {
					res = append(res, i)
					break
				}
			} else {
				break
			}
		}
	}
	return res
}

func copyMap(s map[string]int) map[string]int {
	res := make(map[string]int)
	for k, v := range s {
		res[k] = v
	}
	return res
}

func checkWords(s map[string]int) bool {
	for _, v := range s {
		if v > 0 {
			return false
		}
	}
	return true
}
