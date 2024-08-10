package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取

题目要求出总和为 sum 的所有组合，组合需要去重。
这一题和第 47 题类似，只不过元素可以反复使用。

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var tmp []int
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, &result, tmp)
	return result
}

func findCombinationSum(candidates []int, target int, index int, result *[][]int, tmp []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int, len(tmp)) // 公用一个数组，需要拷贝。把当前切片备份
		copy(b, tmp)
		*result = append(*result, b)
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		tmp = append(tmp, candidates[i])
		findCombinationSum(candidates, target-candidates[i], i, result, tmp)
		tmp = tmp[:len(tmp)-1] // 这步切片
	}
}

func Test_combinationSum(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum1([]int{2, 3, 5}, 8))
}

// dfs
func combinationSum1(candidates []int, target int) [][]int {
	var (
		result [][]int
		dfs    func(i int, combination []int, total int)
	)

	dfs = func(i int, combination []int, total int) {
		if total == target {
			var r = make([]int, len(combination))
			copy(r, combination)
			result = append(result, r)

			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		combination = append(combination, candidates[i])
		dfs(i, combination, total+candidates[i])

		combination = combination[:len(combination)-1]
		dfs(i+1, combination, total)
	}

	dfs(0, []int{}, 0)

	return result
}
