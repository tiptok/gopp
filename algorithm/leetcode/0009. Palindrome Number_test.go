package leetcode

import (
	"math"
	"strconv"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

// 3ms
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 9 {
		return true
	}
	tmp := x
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result == tmp
}

// 8ms
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	if x <= 9 {
		return true
	}
	var tmp = x
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	return result == tmp
}

// 16ms
func isPalindrome1(x int) bool {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return false
	}
	if x < 0 {
		return false
	}
	var tmp = x
	var result int
	for x != 0 {
		result = result*10 + int(x%10)
		x = x / 10
	}
	return result == tmp
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)

	return str == reversePalindrome(str)
}
func reversePalindrome(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
