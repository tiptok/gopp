package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var tmp []int
	sort.Ints(candidates) // 这里是去重的关键逻辑
	findCombinationSum2(candidates, target, 0, &result, tmp)
	return result
}

func findCombinationSum2(candidates []int, target int, index int, result *[][]int, tmp []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int, len(tmp))
		copy(b, tmp)
		*result = append(*result, b)
		return
	}
	for i := index; i < len(candidates); i++ {
		// 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if target >= candidates[i] {
			tmp = append(tmp, candidates[i])
			findCombinationSum2(candidates, target-candidates[i], i+1, result, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
