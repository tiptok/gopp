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

// 46. Permutations
func Test_permute(t *testing.T) {
	input := []int{1, 2, 3}
	excecpt := 6
	out := permute(input)
	if len(out) != excecpt {
		t.Fatalf("out : %v except:%v", len(out), excecpt)
	}
}
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	backtrack_permute(nums, visited, list, &result)
	return result
}

func backtrack_permute(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		//添加元素
		list = append(list, nums[i])
		visited[i] = true
		backtrack_permute(nums, visited, list, result)
		// 移除元素
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}

func Test_permuteUnique(t *testing.T) {
	input := []int{1, 2, 1}
	excecpt := 3
	out := permuteUnique(input)
	if len(out) != excecpt {
		t.Fatalf("out : %v except:%v", len(out), excecpt)
	}
}
func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	visited := make([]bool, len(nums))
	list := make([]int, 0)
	sort.Ints(nums)
	backtrack_permuteUnique(nums, visited, list, &result)
	return result
}

func backtrack_permuteUnique(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// 上一个元素和当前相同，并且没有访问过就跳过
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack_permuteUnique(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
