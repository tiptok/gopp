package leetcode

import (
	"fmt"
	"log"
	"math"
	"testing"
)

//计算范围内 [5,7] 数与的结果
func Test_rangeBitWiseAnd(t *testing.T) {
	input := []struct {
		M int
		N int
	}{
		{M: 0, N: 3},
		{M: 5, N: 7},
	}
	for i := range input {
		out := rangeBitwiseAnd(input[i].M, input[i].N)
		log.Println(fmt.Sprintf("M:%d N:%d Result:%d", input[i].M, input[i].N, out))
	}
}

func rangeBitwiseAnd(m int, n int) int {
	v := math.MaxInt32
	for (m & v) != (n & v) {
		v = v << 1
	}
	return m & v
}
