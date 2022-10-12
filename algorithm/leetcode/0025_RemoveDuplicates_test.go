package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	input := [][]int{
		{0, 0, 1, 1, 2, 2, 3, 4, 4, 5, 5},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{0, 0},
		{1, 1},
		{},
	}
	for i := 0; i < len(input); i++ {
		out := removeDuplicates(input[i])
		PrintResult(out, "", input[i])
	}
}

func removeDuplicates(nums []int) int {
	var (
		cur int = 0
		pre int = 0
	)
	if len(nums) == 0 {
		return 0
	}
	for pre < len(nums) {
		if cur == pre {
			pre++
			continue
		}
		if nums[cur] != nums[pre] {
			cur++
			nums[cur] = nums[pre]
		}
		pre++
	}
	for i := cur + 1; i < len(nums); i++ {
		nums[i] = 0
	}
	return cur + 1
}
