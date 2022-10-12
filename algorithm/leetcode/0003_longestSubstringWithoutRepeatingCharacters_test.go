package leetcode

import (
	"bytes"
	"log"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	inputSlice := []struct {
		Input string
		Want  string
	}{
		{Input: "aaaaa", Want: "a"},
		{Input: "abcabcab", Want: "abc"},
		{Input: "pwwkew", Want: "wke"},
	}
	for _, in := range inputSlice {
		output := longestSubstring(in.Input)
		if in.Want == output {
			log.Printf("Success Input:%v Output:%v Want:%v", in.Input, output, in.Want)
		} else {
			log.Printf("Failer Input:%v Output:%v Want:%v", in.Input, output, in.Want)
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	return len(longestSubstring(s))
}

func longestSubstring(input string) (output string) {
	var (
		lenInput int           = len(input)
		curSub   *bytes.Buffer = bytes.NewBufferString("")
		tmpSub   *bytes.Buffer = bytes.NewBufferString("")
	)
	if lenInput <= 1 {
		return input
	}
	for i := 0; i < lenInput; i++ {
		//检查与tmpSub无重复
		isRepeat, reIndex := checkIsReapet(tmpSub.String(), input[i])
		if !isRepeat {
			tmpSub.WriteString(string(input[i]))
		} else {
			tmpSub.WriteString(string(input[i]))
			tmpStr := tmpSub.String()
			tmpSub.Reset()
			tmpSub.WriteString(tmpStr[reIndex+1:])
		}
		//与curSub比较
		if tmpSub.Len() > curSub.Len() {
			curSub.Reset()
			curSub.WriteString(tmpSub.String())
		}
	}
	return curSub.String()
}

func checkIsReapet(input string, v uint8) (bool, int) {
	for i := 0; i < len(input); i++ {
		if input[i] == v {
			return true, i
		}
	}
	return false, 0
}
