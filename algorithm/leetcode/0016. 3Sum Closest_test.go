package leetcode

import (
	"math"
	"sort"
)

// 给定一个数组，要求在这个数组中找出 3 个数之和离 target 最近
// Given array nums = [-1, 2, 1, -4], and target = 1.
// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2)

func threeSumClosest(nums []int, target int) int {

	// Helper func
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	// Sort the input array in ascending order
	sort.Ints(nums)

	// Initialize the result variable with the sum of the first three elements
	result := nums[0] + nums[1] + nums[2]

	// Iterate through the array
	for i := 0; i < len(nums)-2; i++ {

		// Use two pointers approach
		left := i + 1
		right := len(nums) - 1

		// Continue until the two pointers meet
		for left < right {
			// Calculate the sum of three elements
			sum := nums[i] + nums[left] + nums[right]

			// If the sum is equal to the target, return the sum
			if sum == target {
				return sum
			}

			// Update the result if the current sum is closer to the target
			if abs(target-sum) < abs(target-result) {
				result = sum
			}

			// Move the pointers based on the sum
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// 解法二 暴力解法 O(n^3)
func threeSumClosest1(nums []int, target int) int {
	var result, targetSub = 0, math.MaxInt16
	for i := 0; i < len(nums); i++ {
		for j := i + 1; i < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < targetSub {
					targetSub = abs(nums[i] + nums[j] + nums[k] - target)
					result = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
