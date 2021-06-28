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
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if dp[i-1][0] == 1 && obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		}
	}
	for j := 1; j < n; j++ {
		if dp[0][j-1] == 1 && obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

/*序列类型*/
func Test_climbStairs(t *testing.T) {
	input := 3
	excecpt := 3
	out := climbStairs(input)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}

// 70.climb stairs
func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 55. Jump Game
func Test_canJump(t *testing.T) {
	input := []int{3, 2, 1, 0, 4} //{2, 3, 1, 1, 4}
	excecpt := false
	out := canJump2(input)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}

// 70.climb stairs
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
			}
		}
	}
	return f[len(nums)-1]
}

func canJump2(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var reach = nums[0]
	i := 0
	for i = 1; i < len(nums) && i <= reach; i++ {
		reach = max(i+nums[i], reach)
	}
	return i == len(nums)
}

// 132.Palindrome Partitioning II  ***
func Test_minCut(t *testing.T) {
	input := "aab"
	excecpt := 1
	out := minCut(input)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}

func minCut(s string) int {
	for len(s) == 0 || len(s) == 1 {
		return 0
	}
	f := make([]int, len(s)+1)
	f[0] = -1
	f[1] = 0
	for i := 1; i <= len(s); i++ {
		f[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(s)]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 300. Longest Increasing Subsequence
func Test_lengthOfLIS(t *testing.T) {
	input := []int{10, 9, 2, 5, 3, 7, 101, 18}
	excecpt := 4
	out := lengthOfLIS(input)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	result := f[0]
	for i := 1; i < len(nums); i++ {
		result = max(result, f[i])
	}
	return result
}

// 139. Longest Increasing Subsequence
func Test_wordBreak(t *testing.T) {
	input1 := "bb"
	input2 := []string{"a", "b", "bbb", "bbbb"}
	excecpt := true
	out := wordBreak(input1, input2)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	dict := make(map[string]bool, len(wordDict))
	val := false
	for i := 0; i < len(wordDict); i++ {
		dict[wordDict[i]] = val
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			word := string(s[j:i])
			if _, ok := dict[word]; f[j] && ok {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

/***Two Sequences DP***/

// 139. Longest Increasing Subsequence
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）
// 后组成的新字符串。 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列
func Test_longestCommonSubsequence(t *testing.T) {
	table := []struct {
		input1 string
		input2 string
		except int
	}{
		// {
		// 	"psnw",
		// 	"vozsh",
		// 	1,
		// },
		{
			"ezupkr",
			"ubmrapg",
			2,
		},
	}
	for i := range table {
		out := longestCommonSubsequence(table[i].input1, table[i].input2)
		if out != table[i].except {
			t.Fatalf("out : %v except:%v", out, table[i].except)
		}
	}
}
func longestCommonSubsequence(a string, b string) int {
	// dp[i][j] a前i个和b前j个字符最长公共子序列
	// dp[m+1][n+1]
	//   ' a d c e
	// ' 0 0 0 0 0
	// a 0 1 1 1 1
	// c 0 1 1 2 1
	//
	f := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		f[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[len(a)][len(b)]
}

// 72. Edit Distance
func Test_minDistance(t *testing.T) {
	table := []struct {
		input1 string
		input2 string
		except int
	}{
		{
			"horse",
			"ros",
			3,
		},
	}
	for i := range table {
		out := minDistance(table[i].input1, table[i].input2)
		if out != table[i].except {
			t.Fatalf("out : %v except:%v", out, table[i].except)
		}
	}
}
func minDistance(a string, b string) int {
	f := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		f[i] = make([]int, len(b)+1)
	}
	for i := 0; i < len(f); i++ {
		f[i][0] = i
	}
	for j := 0; j < len(f[0]); j++ {
		f[0][j] = j
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				// 否则取删除、插入、替换最小操作次数的值+1
				f[i][j] = min(min(f[i-1][j], f[i][j-1]), f[i-1][j-1]) + 1
			}
		}
	}
	return f[len(a)][len(b)]
}

/***零钱和背包***/

// 322.Coin Change  ###				
//给定不同面额的硬币 coins 和一个总金额 amount。
//编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
//如果没有任何一种硬币组合能组成总金额，返回 -1
func Test_coinChange(t *testing.T) {
	input1 := []int{1, 2, 5}
	input2 := 11
	excecpt := 3
	out := coinChange(input1, input2)
	if out != excecpt {
		t.Fatalf("out : %v except:%v", out, excecpt)
	}
}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
