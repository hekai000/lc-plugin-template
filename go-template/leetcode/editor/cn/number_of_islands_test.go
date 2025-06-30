/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=30104
 *
 * [200] 岛屿数量
 */

package leetcode_solutions

import "testing"

// @lc code=start
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfsni(grid, i, j)
			}
		}
	}
	return res
}

func dfsni(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2'
	dfsni(grid, r+1, c)
	dfsni(grid, r-1, c)
	dfsni(grid, r, c+1)
	dfsni(grid, r, c-1)
	
}

// @lc code=end

func TestNumberOfIslands(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [\n["1","1","1","1","0"],\n["1","1","0","1","0"],\n["1","1","0","0","0"],\n["0","0","0","0","0"]\n]\n
// @lcpr case=end

// @lcpr case=start
// [\n["1","1","0","0","0"],\n["1","1","0","0","0"],\n["0","0","1","0","0"],\n["0","0","0","1","1"]\n]\n
// @lcpr case=end

*/
