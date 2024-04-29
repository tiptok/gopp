package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1+num2))
	byteToInt := func(num byte) int {
		return int(num - '0')
	}
	b1 := []byte(num1)
	b2 := []byte(num2)
	for j := 0; j < len(b2); j++ {
		for i := 0; i < len(b1); i++ {
			val := byteToInt(b1[i]) * byteToInt(b2[j])
			tmp[i+j+1] += val
		}
	}
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	if tmp[0] == 0 {
		tmp = tmp[1:]
	}
	res := make([]byte, len(tmp))
	for i := 0; i < len(tmp); i++ {
		res[i] = '0' + byte(tmp[i])
	}
	return string(res)
}

func TestMultiply(t *testing.T) {
	//assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "56088", multiply("123", "456"))
}
