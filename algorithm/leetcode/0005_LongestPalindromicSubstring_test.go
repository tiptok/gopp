package leetcode

import (
	"fmt"
	"log"
	"testing"
)

/*
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
*/

func TestLongestPalindrome(t *testing.T) {
	input := []string{
		"abccbabc",
		"aba",
	}
	for _, in := range input {
		out := longestPalindrome(in)
		log.Println(fmt.Sprintf("Input:%s OutPut:%s", in, out))
	}
}

// 动态规划
// https://leetcode.com/problems/longest-palindromic-substring/discuss/2921/Share-my-Java-solution-using-dynamic-programming
//func longestPalindrome(s string) string {
//	n := len(s)
//	dp := make([][]bool, n)
//	res := ""
//	for i := n - 1; i >= 0; i-- {
//		dp[i] = make([]bool, n)
//		for j := i; j < n; j++ {
//			dp[i][j] = (s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]))
//			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
//				res = s[i : j+1]
//			}
//		}
//	}
//	return res
//}

// 中心扩散法
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
