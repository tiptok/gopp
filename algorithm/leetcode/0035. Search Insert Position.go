package leetcode

func searchInsert(nums []int, target int) int {
	var i = 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}
