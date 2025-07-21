/*
 * @lc app=leetcode.cn id=994 lang=golang
 * @lcpr version=30104
 *
 * [994] 腐烂的橘子
 */

package leetcode_solutions

import "testing"

// @lc code=start
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	step := 0

	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x, y := cur[0]+dir[0], cur[1]+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					queue = append(queue, []int{x, y})
				}
			}

		}
		step++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if step == 0 {
		return 0
	}
	return step - 1

}

// @lc code=end

func TestRottingOranges(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[2,1,1],[1,1,0],[0,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,1],[0,1,1],[1,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,2]]\n
// @lcpr case=end

*/
