package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	input := " 0000000000012345678"
	out := myAtoi(input)
	assert.Equal(t, 82, out)
}

func myAtoi(s string) int {
	maxInt := int(2 << 30)
	signAllowed := true
	whitespaceAllowed := true
	sign := 1
	digits := []int{}
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
			if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		if c < '0' || c > '9' {
			break
		}
		signAllowed = false
		whitespaceAllowed = false
		digits = append(digits, int(c)-int('0'))
	}

	var num, place int = 0, 1
	// 去掉前导的0 例如“000123”
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}

	for i := len(digits) - 1; i >= 0; i-- {
		num += digits[i] * place
		place *= 10
	}
	if len(digits) > 10 {
		if sign > 0 {
			return maxInt - 1
		} else {
			return -maxInt
		}
	}

	num *= sign
	if -maxInt > num || len(digits) > 10 {
		return -maxInt
	}
	if maxInt <= num || len(digits) > 10 {
		return maxInt - 1
	}
	return num
}
