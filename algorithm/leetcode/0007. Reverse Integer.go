package leetcode

import "math"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
注意:假设我们的环境只能存储得下 32 位的有符号整数，
则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + int(x%10)
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func reverse1(x int) int {
	//if x > math.MaxInt32 || x < math.MinInt32 {
	//	return 0
	//}
	var isNegative = x < 0
	if isNegative {
		x = -x
	}
	var list []byte
	for x > 0 {
		list = append(list, byte(x%10))
		x = x / 10
	}
	var result int
	for _, val := range list {
		if result == 0 && val == 0 {
			continue
		}
		result = result*10 + int(val)
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if isNegative {
		return -result
	}
	return result
}
