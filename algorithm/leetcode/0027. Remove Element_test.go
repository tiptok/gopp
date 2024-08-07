package leetcode

/*

给定一个数组 nums 和一个数值 val，将数组中所有等于 val 的元素删除，并返回剩余的元素个数。

解题思路 #
这道题和第 283 题很像。这道题和第 283 题基本一致，283 题是删除 0，这一题是给定的一个 val，实质是一样的。

这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数，OJ 最终判断题目的时候会读取数组剩余个数的元素进行输出。

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2
*/

func removeElement(nums []int, val int) int {
	var j int
	for i, target := range nums {
		if target != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
