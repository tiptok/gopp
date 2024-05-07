package leetcode

import "strings"

/*
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/

func numDecodings1(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

var m = map[string]string{
	"1":  "A",
	"2":  "B",
	"3":  "C",
	"4":  "D",
	"5":  "E",
	"6":  "F",
	"7":  "G",
	"8":  "H",
	"9":  "I",
	"10": "J",
	"11": "K",
	"12": "L",
	"13": "M",
	"14": "N",
	"15": "O",
	"16": "P",
	"17": "Q",
	"18": "R",
	"19": "S",
	"20": "T",
	"21": "U",
	"22": "V",
	"23": "W",
	"24": "Z",
	"25": "Y",
	"26": "Z",
}

func numDecodings(s string) int {
	cache := make(map[string]int)
	return numDecodingsCached(s, cache)
}

func numDecodingsCached(s string, cache map[string]int) int {
	if len(s) == 0 {
		return 1
	} else if v, ok := cache[s]; ok {
		return v
	}
	var n = 0
	for k, _ := range m {
		if strings.HasPrefix(s, k) {
			ss := strings.TrimPrefix(s, k)
			n += numDecodingsCached(ss, cache)
		}
	}

	cache[s] = n
	return n
}
