package code

import (
	"fmt"
	"sort"
	"testing"
)

/*回溯法*/

// 78.subsets
func Test_subsets(t *testing.T) {
	input := []int{1, 2, 3}
	excecpt := 8
	out := subsets(input)
	if len(out) != excecpt {
		t.Fatalf("out : %v except:%v", len(out), excecpt)
	}
}
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tmpList := make([]int, 0)
	backtrack(nums, 0, tmpList, &res)
	return res
}

func backtrack(nums []int, start int, tmpList []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	copyList := make([]int, len(tmpList))
	copy(copyList, tmpList)
	*result = append(*result, copyList)
	fmt.Println(start, copyList)
	for i := start; i < len(nums); i++ {
		tmpList = append(tmpList, nums[i])
		backtrack(nums, i+1, tmpList, result)
		// 撤销
		tmpList = tmpList[0 : len(tmpList)-1]
	}
}

//90. Subsets II
func Test_subsetsWithDup(t *testing.T) {
	input := []int{1, 2, 2}
	excecpt := 6
	out := subsetsWithDup(input)
	if len(out) != excecpt {
		t.Fatalf("out : %v except:%v", len(out), excecpt)
	}
}
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	tmpList := make([]int, 0)
	sort.Ints(nums)
	backtrackII(nums, 0, tmpList, &res)
	return res
}

func backtrackII(nums []int, start int, tmpList []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	copyList := make([]int, len(tmpList))
	copy(copyList, tmpList)
	*result = append(*result, copyList)
	fmt.Println(start, copyList)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tmpList = append(tmpList, nums[i])
		backtrackII(nums, i+1, tmpList, result)
		// 撤销
		tmpList = tmpList[0 : len(tmpList)-1]
	}
}
