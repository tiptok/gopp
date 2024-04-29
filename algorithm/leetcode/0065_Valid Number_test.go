package leetcode

/*
For example, all the following are valid numbers: ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3",
"3e+7", "+6e-1", "53.5e93", "-123.456e789"],
while the following are not valid numbers:
["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].

用三个变量分别标记是否出现过数字、是否出现过’.‘和 是否出现过 ‘e/E’
从左到右依次遍历字符串中的每一个元素
如果是数字，则标记数字出现过
如果是 ‘.’, 则需要 ‘.‘没有出现过，并且 ‘e/E’ 没有出现过，才会进行标记
如果是 ‘e/E’, 则需要 ‘e/E’没有出现过，并且前面出现过数字，才会进行标记
如果是 ‘+/-’, 则需要是第一个字符，或者前一个字符是 ‘e/E’，才会进行标记，并重置数字出现的标识
最后返回时，需要字符串中至少出现过数字，避免下列case: s == ‘.’ or ‘e/E’ or ‘+/e’ and etc…
*/

func isNumber(s string) bool {
	numFlag, dotFlag, eFlag := false, false, false
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			numFlag = true
		} else if s[i] == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if (s[i] == 'e' || s[i] == 'E') && numFlag && !eFlag {
			eFlag = true
			numFlag = false // reJudge integer after 'e' or 'E'
		} else if (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'e' || s[i-1] == 'E') {
			continue
		} else {
			return false
		}
	}
	return numFlag
}
