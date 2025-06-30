/*
 * @lc app=leetcode.cn id=1020 lang=golang
 * @lcpr version=30104
 *
 * [1020] 飞地的数量
 */

package leetcode_solutions

import "testing"

// @lc code=start
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		dfsev(grid, i, 0)
		dfsev(grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfsev(grid, 0, j)
		dfsev(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++

			}
		}
	}
	return res
}

func dfsev(grid [][]int, r, c int) {
	m, n := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if grid[r][c] == 0 || grid[r][c] == 2 {
		return
	}
	grid[r][c] = 2

	dfsev(grid, r+1, c)
	dfsev(grid, r-1, c)
	dfsev(grid, r, c+1)
	dfsev(grid, r, c-1)
}

// @lc code=end

func TestNumberOfEnclaves(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]\n
// @lcpr case=end

*/
