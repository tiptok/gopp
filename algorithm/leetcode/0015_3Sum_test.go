package leetcode

import (
	"log"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := [][]int{
		[]int{-1, -1, 0, 1, 2, -1, -4},
		[]int{0, 0},
		[]int{0, 0, 0}, //isDo
		[]int{-2, -2, 4},
		[]int{-1, 0, 0, 1},
		[]int{-5, -4, -1, -1, -1, 0, 1, 2, 2, 4},
	}
	for _, in := range input {
		out := threeSum(in)
		log.Printf("In:%v Out:%v \n\n", in, out)
	}
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	retData := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		left, right := i, len(nums)-1
		for left < right {
			sum := a + nums[left] + nums[right]
			if left == i {
				left++
			} else if right == i {
				right--
			} else if sum == 0 {
				tmpData := []int{a, nums[left], nums[right]}
				//sort.Ints(tmpData)
				retData = append(retData, tmpData)
				for {
					if nums[left] == nums[left+1] {
						left++
					} else {
						break
					}
					if left >= len(nums)-1 {
						break
					}
				}
				left++
				for {
					if nums[right] == nums[right-1] {
						right--
					} else {
						break
					}
					if right <= 0 {
						break
					}
				}
				right--
				continue
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return retData
}
