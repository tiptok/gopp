package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	result, last, tmp := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i : i+1]
		tmp = roman[c]
		if tmp < last {
			result -= tmp
		} else {
			result += tmp
		}
		last = tmp
	}
	return result
}

func Test_romanToInt(t *testing.T) {
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}

func romanToInt1(s string) int {
	var result = 0
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(s); i++ {
		j := 0
		sub1 := s[i : i+1]
		sub2 := ""
		if i <= len(s)-2 {
			sub2 = s[i : i+2]
		}
		for j < len(symbols) {
			if sub1 == symbols[j] {
				break
			}
			if sub2 == symbols[j] {
				i++
				break
			}
			j++
		}
		result += values[j]
	}
	return result
}
