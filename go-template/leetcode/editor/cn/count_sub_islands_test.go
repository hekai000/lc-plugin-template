/*
 * @lc app=leetcode.cn id=1905 lang=golang
 * @lcpr version=30104
 *
 * [1905] 统计子岛屿
 */

package leetcode_solutions

import "testing"

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfscsi(grid2, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfscsi(grid2, i, j)
			}
		}
	}
	return res
}

func dfscsi(grid [][]int, r, c int) {
	m, n := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if grid[r][c] == 0 || grid[r][c] == 2 {
		return
	}
	grid[r][c] = 2

	dfscsi(grid, r+1, c)
	dfscsi(grid, r-1, c)
	dfscsi(grid, r, c+1)
	dfscsi(grid, r, c-1)
}

// @lc code=end

func TestCountSubIslands(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]]\n\n[[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]]\n\n[[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]\n
// @lcpr case=end

*/
