package code

import (
	"testing"
)

// 动态规划 DP

/************* 120 . Triangle *********/
func Test_minimumTotal(t *testing.T) {
	input := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	out := minimumTotalTravelDes2(input)
	if out != 11 {
		t.Fatal("fatal", out)
	}
}

// 方法一: DFS + HASH
var hashDFS = make(map[int]map[int]int)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	x, y := 0, 0
	for i := range triangle {
		hashDFS[i] = make(map[int]int)
	}
	min := dfsMinimumTotal(triangle, x, y)
	return min
}

func dfsMinimumTotal(triangle [][]int, x, y int) int {
	if x == len(triangle) {
		return 0
	}
	if v, ok := hashDFS[x][y]; ok {
		return v
	}
	hashDFS[x][y] = min(dfsMinimumTotal(triangle, x+1, y), dfsMinimumTotal(triangle, x+1, y+1)) + triangle[x][y]

	return hashDFS[x][y]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法二:遍历 底部到顶部

func minimumTotalTravelDes(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	var hashDFS = make(map[int]map[int]int)
	for i := range triangle {
		if hashDFS[i] == nil {
			hashDFS[i] = make(map[int]int)
		}
		for j := 0; j < len(triangle[i]); j++ {
			hashDFS[i][j] = triangle[i][j]
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			hashDFS[i][j] = min(hashDFS[i+1][j], hashDFS[i+1][j+1]) + triangle[i][j]
		}
	}
	return hashDFS[0][0]
}

func minimumTotalTravelDes2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// 复用triangle , 优化空间
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

// 方法二:遍历 顶部到底部
func minimumTotalTravelDes3(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	return 0
}

// 128. Longest Consecutive (不属于动态规划)
func Test_longestConsecutive(t *testing.T) {
	input := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1} //{100,4,200,1,2,3}
	out := longestConsecutive(input)
	if out != 9 {
		t.Fatalf("out : %v except:%v", out, 9)
	}
}

func longestConsecutive(nums []int) int {
	var res = 0
	var s = make(map[int]int)
	for i := range nums {
		num := nums[i]
		if _, ok := s[num]; !ok {
			s[num] = 1
		}
		// else{
		// 	s[num] ++
		// }
	}

	for i := 0; i < len(nums); i++ {
		var count = 1
		num := nums[i]
		// if v,ok:=s[num];ok{
		// 	count = v
		// }
		num--
		v, ok := s[num]
		for ok {
			num--
			count += v
			v, ok = s[num]
			delete(s, num)
		}

		num = nums[i]
		num++
		v, ok = s[num]
		for ok {
			num++
			count += v
			v, ok = s[num]
			delete(s, num)
		}
		res = max(count, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
