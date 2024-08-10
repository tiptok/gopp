package leetcode

/*
找到缺失的第一个正整数。

为了减少时间复杂度，可以把 input 数组都装到 map 中，然后 i 循环从 1 开始，
依次比对 map 中是否存在 i，只要不存在 i 就立即返回结果，即所求。
*/
func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v] = v
	}
	for index := 1; index < len(nums)+1; index++ {
		if _, ok := numMap[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}
