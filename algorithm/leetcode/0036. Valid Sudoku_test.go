package leetcode

import "strconv"

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
// 暴力求解，有更优解
func isValidSudoku(board [][]byte) bool {
	// 判断行
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := string(board[i][j])
			if cellVal == "." {
				continue
			}
			index, _ := strconv.Atoi(cellVal)
			if index > 9 || index < 1 {
				return false
			}
			if tmp[index] == 1 {
				return false
			}
			tmp[index] = 1
		}
	}

	// 判断列
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := string(board[j][i])
			if cellVal == "." {
				continue
			}
			index, _ := strconv.Atoi(cellVal)
			if index > 9 || index < 1 {
				return false
			}
			if tmp[index] == 1 {
				return false
			}
			tmp[index] = 1
		}
	}
	// 判断 9宫格 3X3 cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}

	return true
}
