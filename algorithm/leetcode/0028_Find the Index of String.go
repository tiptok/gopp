package leetcode

/*
Input: haystack = "hello", needle = "ll"
Output: 2

实现一个查找 substring 的函数。如果在母串中找到了子串，返回子串在母串中出现的下标，如果没有找到，返回 -1，如果子串是空串，则返回 0 。
*/

func strStr(haystack string, needle string) int {
	var needleLen = len(needle)
	if len(haystack) < needleLen {
		return -1
	}
	for i := 0; i <= len(haystack)-needleLen; i++ {
		if string(haystack[i:i+needleLen]) == needle {
			return i
		}
	}
	return -1
}
