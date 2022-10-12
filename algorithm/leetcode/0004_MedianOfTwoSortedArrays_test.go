package leetcode

import (
	"log"
	"sort"
	"testing"
)

func TestFindMedian(t *testing.T) {
	input := []struct {
		Num1 []int
		Num2 []int
		Want float64
	}{
		{Num1: []int{1, 3}, Num2: []int{2}, Want: 2.0},
		{Num1: []int{1, 2}, Num2: []int{3, 4}, Want: 2.5},
		{Num1: []int{1, 2, 7, 8}, Num2: []int{3, 4, 5}, Want: 4.0},
		{Num1: []int{}, Num2: []int{3}, Want: 3.0},
	}
	for _, in := range input {
		want := findMedianSortedArrays(in.Num1, in.Num2)
		if want == in.Want {
			log.Printf("Success Num1:%v Num2:%v Output:%v Want:%v", in.Num1, in.Num2, want, in.Want)
		} else {
			log.Printf("Failer Num1:%v Num2:%v Output:%v Want:%v", in.Num1, in.Num2, want, in.Want)
		}
	}
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMidian(nums2)
	}
	if len(nums2) == 0 {
		return getMidian(nums1)
	}
	nums := make([]int, 0)
	nums = append(nums1, nums2...)
	sort.Ints(nums)
	return getMidian(nums)
}

func getMidian(nums []int) float64 {
	log.Printf("Nums:%v\n", nums)
	total := len(nums)
	mid1, mid2 := 0, 0
	if total%2 > 0 {
		mid1 = nums[total/2]
		mid2 = nums[total/2]
	} else {
		i := total / 2
		mid1 = nums[i-1]
		mid2 = nums[i]
	}
	return float64(mid1+mid2) / 2.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMidian(nums2)
	}
	if len(nums2) == 0 {
		return getMidian(nums1)
	}
	total := len(nums1) + len(nums2)
	nums := make([]int, total)
	i1, i2 := 0, 0
	for i := 0; i < total; i++ {
		nv1, nv2 := 0, 0
		if i1 < len(nums1) {
			nv1 = nums1[i1]
		}
		if i2 < len(nums2) {
			nv2 = nums2[i2]
		}
		if nv1 < nv2 && i1 < len(nums1) || i2 == len(nums2) {
			nums[i] = nv1
			i1++
		} else if i2 < len(nums2) {
			nums[i] = nv2
			i2++
		}
	}
	return getMidian(nums)
}
