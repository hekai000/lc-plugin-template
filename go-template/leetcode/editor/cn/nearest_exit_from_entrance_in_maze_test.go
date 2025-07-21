/*
 * @lc app=leetcode.cn id=1926 lang=golang
 * @lcpr version=30104
 *
 * [1926] 迷宫中离入口最近的出口
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	queue := [][]int{entrance}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[entrance[0]][entrance[1]] = true
	step := 0
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if isExit(maze, entrance, cur[0], cur[1]) {
				return step
			}
			for _, v := range getNeighborsNE(maze, cur) {
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

func getNeighborsNE(maze [][]byte, cur []int) [][]int {
	neighbors := [][]int{}
	m, n := len(maze), len(maze[0])
	row, col := cur[0], cur[1]
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range dirs {
		newR := row + dir[0]
		newC := col + dir[1]
		if newR < 0 || newR >= m || newC < 0 || newC >= n || maze[newR][newC] == '+' {
			continue
		}

		neighbors = append(neighbors, []int{newR, newC})

	}
	return neighbors
}
func isExit(maze [][]byte, entrance []int, i, j int) bool {
	m, n := len(maze), len(maze[0])
	if maze[i][j] == '.' && !(i == entrance[0] && j == entrance[1]) {
		if i == 0 || j == 0 || i == m-1 || j == n-1 {
			return true
		}
	}
	return false
}

// @lc code=end

func TestNearestExitFromEntranceInMaze(t *testing.T) {
	// your test code here
	maze := [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}
	entrance := []int{1, 2}
	fmt.Println(nearestExit(maze, entrance))

}

/*
// @lcpr case=start
// [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]]\n[1,2]\n
// @lcpr case=end

// @lcpr case=start
// [["+","+","+"],[".",".","."],["+","+","+"]]\n[1,0]\n
// @lcpr case=end

// @lcpr case=start
// [[".","+"]]\n[0,0]\n
// @lcpr case=end

*/
