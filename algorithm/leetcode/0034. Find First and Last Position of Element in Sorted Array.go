package leetcode

func searchRange(nums []int, target int) []int {
	var i = 0
	var j = 0
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	for ; i < len(nums); i++ {
		if nums[i] == target {
			break
		}
	}
	for j = len(nums) - 1; j >= i; j-- {
		if nums[j] == target {
			break
		}
	}
	if j < i {
		return []int{-1, -1}
	}
	return []int{i, j}
}
