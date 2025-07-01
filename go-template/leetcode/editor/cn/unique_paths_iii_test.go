/*
 * @lc app=leetcode.cn id=980 lang=golang
 * @lcpr version=30104
 *
 * [980] 不同路径 III
 */

package leetcode_solutions

import "testing"

// @lc code=start
func uniquePathsIII(grid [][]int) int {
	result := 0
	var visitedCount int
	m, n := len(grid), len(grid[0])
	var total, startr, startc int
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startr, startc = i, j
			}
			if grid[i][j] == 1 || grid[i][j] == 0 {
				total++
			}
		}
	}
	dfsup(grid, startr, startc, &result, &visitedCount, total, visited, dirs)
	return result

}

func dfsup(grid [][]int, r int, c int, res *int, visitedCount *int, total int, visited [][]bool, dirs [][]int) {
	m, n := len(grid), len(grid[0])
	if r < 0 || c < 0 || r >= m || c >= n {
		return
	}
	if grid[r][c] == -1 || visited[r][c] {
		return
	}

	if grid[r][c] == 2 {
		if *visitedCount == total {
			*res++
		}
		return
	}
	visited[r][c] = true
	*visitedCount++
	for _, dir := range dirs {
		dfsup(grid, r+dir[0], c+dir[1], res, visitedCount, total, visited, dirs)
	}
	visited[r][c] = false
	*visitedCount--
}

// @lc code=end

func TestUniquePathsIii(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0,0],[0,0,0,0],[0,0,0,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[2,0]]\n
// @lcpr case=end

*/
