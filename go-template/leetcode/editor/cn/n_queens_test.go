/*
 * @lc app=leetcode.cn id=51 lang=golang
 * @lcpr version=30104
 *
 * [51] N 皇后
 */

package leetcode_solutions

import (
	"strings"
	"testing"
)

// @lc code=start
func solveNQueens(n int) [][]string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	res := [][]string{}
	var backtrackQ func(row int)
	backtrackQ = func(row int) {
		if row == len(board) {
			temp := make([]string, len(board))
			copy(temp, board)
			res = append(res, temp)
			return

		}
		for col := 0; col < n; col++ {
			if !isValidsnq(board, row, col) {
				continue
			}
			rowChars := []rune(board[row])
			rowChars[col] = 'Q'
			board[row] = string(rowChars)
			backtrackQ(row + 1)
			rowChars[col] = '.'
			board[row] = string(rowChars)
		}
	}
	backtrackQ(0)
	return res
}

func isValidsnq(board []string, row, col int) bool {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true

}

// @lc code=end

func TestNQueens(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
