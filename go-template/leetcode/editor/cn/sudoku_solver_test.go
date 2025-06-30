/*
 * @lc app=leetcode.cn id=37 lang=golang
 * @lcpr version=30104
 *
 * [37] 解数独
 */

package leetcode_solutions

import "testing"

// @lc code=start
func solveSudoku(board [][]byte) {
	found := false

	var backtracksd func(index int)
	backtracksd = func(index int) {
		if found {
			return
		}
		if index == 9*9 {
			found = true
			return
		}

		r, c := index/9, index%9
		if board[r][c] != '.' {
			backtracksd(index + 1)
			return
		}

		for ch := byte('1'); ch <= byte('9'); ch++ {
			if !isValid(ch, r, c, board) {
				continue
			}
			board[r][c] = ch
			backtracksd(index + 1)
			if found {
				return
			}
			board[r][c] = '.'

		}

	}
	backtracksd(0)
}

func isValid(num byte, r, c int, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][c] == num {
			return false
		}
		if board[r][i] == num {
			return false
		}
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == num {
			return false
		}
	}
	return true

}

// @lc code=end

func TestSudokuSolver(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// \n[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]\n
// @lcpr case=end

*/
