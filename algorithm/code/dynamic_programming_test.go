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
func Test_minPathSum(t *testing.T) {
	input := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}} //{100,4,200,1,2,3}
	out := minPathSum(input)
	if out != 7 {
		t.Fatalf("out : %v except:%v", out, 7)
	}
}

// 64.Minimum Path Sum
// TODO:40 ms, faster than 5.48%
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

// 62. Unique Paths
func Test_uniquePaths(t *testing.T) {
	input := struct {
		m int
		n int
	}{
		m: 3,
		n: 2,
	}
	out := uniquePaths(input.m, input.n)
	if out != 3 {
		t.Fatalf("out : %v except:%v", out, 3)
	}
}
func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}} //{100,4,200,1,2,3}
	out := uniquePathsWithObstacles(input)
	if out != 2 {
		t.Fatalf("out : %v except:%v", out, 2)
	}
}
// 63. Unique Paths II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	if obstacleGrid[0][0]==1{
		return 0
	}
	var dp = make([][]int,m)
	for i:=0; i<m;i++{
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if (dp[i-1][0]==1 && obstacleGrid[i][0]==0){
			dp[i][0] = 1
		}
	}
	for j := 1; j < n; j++ {
		if (dp[0][j-1]==1 && obstacleGrid[0][j]==0){
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j <n; j++ {
			if obstacleGrid[i][j]==1{
				dp[i][j]=0
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}