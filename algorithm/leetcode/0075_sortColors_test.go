package leetcode

import (
	"fmt"
	"log"
	"testing"
)

func Test_sortColors(t *testing.T) {
	input := [][]int{
		{2, 2, 1, 0, 1, 0, 0},
	}
	for i := 0; i < len(input); i++ {
		log.Println(fmt.Sprintf("Test>>> Index:%d", i))
		log.Println(fmt.Sprintf("In:%v", input[i]))
		sortColors(input[i])
		log.Println(fmt.Sprintf("Out:%v", input[i]))
	}
}

func sortColors(nums []int) {
	var (
		begin int = -1
		end   int = len(nums)
	)
	exchange := func(i, j int) {
		k := nums[i]
		nums[i] = nums[j]
		nums[j] = k
	}
	for i := 0; i < end; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			end--
			exchange(i, end)
		} else {
			begin++
			exchange(begin, i)
			i++
		}
	}
}
