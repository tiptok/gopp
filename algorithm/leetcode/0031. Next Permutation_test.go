package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
题目有 3 个问题需要解决。如何找到下一个排列。不存在下一个排列的时候如何生成最小的排列。如何原地修改。
先解决第一个问题，如何找到下一个排列。下一个排列是找到一个大于当前排序的字典序，且变大的幅度最小。
那么只能将较小的数与较大数做一次原地交换。并且较小数的下标要尽量靠右，较大数也要尽可能小。
原地交换以后，还需要将较大数右边的数按照升序重新排列。这样交换以后，才能生成下一个排列。
以排列 [8,9,6,10,7,2] 为例：能找到的符合条件的一对「较小数」与「较大数」的组合为 6 与 7，满足「较小数」尽量靠右，而「较大数」尽可能小。
当完成交换后排列变为 [8,9,7,10,6,2]，此时我们可以重排「较小数」右边的序列，序列变为 [8,9,7,2,6,10]
*/
func nextPermutation(nums []int) {
	// 找最近的i
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// 找最近的j
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverseSwap(&nums, i+1, len(nums)-1)
}

func reverseSwap(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func Test_nextPermutation(t *testing.T) {
	input := []int{8, 9, 6, 10, 7, 2}
	nextPermutation(input)
	assert.Equal(t, []int{8, 9, 7, 2, 6, 10}, input)
}
