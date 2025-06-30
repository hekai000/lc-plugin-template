/*
 * @lc app=leetcode.cn id=695 lang=golang
 * @lcpr version=30104
 *
 * [695] 岛屿的最大面积
 */

package leetcode_solutions

import "testing"

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = maxmoi(res, dfsmoi(grid, i, j))
			}
		}
	}
	return res
}

func dfsmoi(grid [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return 0
	}
	if grid[r][c] == 0 || grid[r][c] == 2 {
		return 0
	}
	grid[r][c] = 2

	return dfsmoi(grid, r+1, c) +
		dfsmoi(grid, r-1, c) +
		dfsmoi(grid, r, c+1) +
		dfsmoi(grid, r, c-1) + 1

}

func maxmoi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestMaxAreaOfIsland(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// \n[[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0,0,0,0,0,0]]\n
// @lcpr case=end

*/
