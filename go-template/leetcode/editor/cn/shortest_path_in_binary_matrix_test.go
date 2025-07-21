/*
 * @lc app=leetcode.cn id=1091 lang=golang
 * @lcpr version=30104
 *
 * [1091] 二进制矩阵中的最短路径
 */

package leetcode_solutions

import "testing"

// @lc code=start
func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	start := []int{0, 0}
	end := []int{len(grid) - 1, len(grid[0]) - 1}
	if grid[0][0] != 0 {
		return -1
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	queue := [][]int{start}
	step := 1
	visited[start[0]][start[1]] = true
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == end[0] && cur[1] == end[1] {
				return step
			}
			for _, v := range getNeighborsSPBMT(grid, m, n, cur[0], cur[1]) {
				if !visited[v[0]][v[1]] {
					queue = append(queue, v)
					visited[v[0]][v[1]] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNeighborsSPBMT(grid [][]int, m int, n int, row int, col int) [][]int {
	neighbors := [][]int{}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}}
	for _, dir := range dirs {
		newR := row + dir[0]
		newC := col + dir[1]
		if newR < 0 || newR >= m || newC < 0 || newC >= n || grid[newR][newC] != 0 {
			continue
		}
		neighbors = append(neighbors, []int{newR, newC})
	}
	return neighbors
}

// @lc code=end

func TestShortestPathInBinaryMatrix(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[0,1],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

*/
