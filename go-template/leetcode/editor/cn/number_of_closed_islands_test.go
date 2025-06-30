/*
 * @lc app=leetcode.cn id=1254 lang=golang
 * @lcpr version=30104
 *
 * [1254] 统计封闭岛屿的数目
 */

package leetcode_solutions

import "testing"

// @lc code=start
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		dfsci(grid, i, 0)
		dfsci(grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfsci(grid, 0, j)
		dfsci(grid, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfsci(grid, i, j)
			}
		}
	}
	return res
}

func dfsci(grid [][]int, r, c int) {
	m, n := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if grid[r][c] == 1 || grid[r][c] == 2 {
		return
	}
	grid[r][c] = 2

	dfsci(grid, r+1, c)
	dfsci(grid, r-1, c)
	dfsci(grid, r, c+1)
	dfsci(grid, r, c-1)
}

// @lc code=end

func TestNumberOfClosedIslands(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1,1,1,1,1],\n[1,0,0,0,0,0,1],\n[1,0,1,1,1,0,1],\n[1,0,1,0,1,0,1],\n[1,0,1,1,1,0,1],\n[1,0,0,0,0,0,1],\n[1,1,1,1,1,1,1]]\n
// @lcpr case=end

*/
