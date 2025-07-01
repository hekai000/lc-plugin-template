/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=30104
 *
 * [79] 单词搜索
 */

package leetcode_solutions

import "testing"

// @lc code=start
func exist(board [][]byte, word string) bool {
	found := true
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfseee func([][]byte, string, int, int, [][]bool, int)
	dfseee = func(board [][]byte, word string, i int, j int, visited [][]bool, pos int) {
		if pos == len(word) {
			found = true
			return
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if visited[i][j] {
			return
		}
		if found {
			return
		}

		if board[i][j] != word[pos] {
			return
		}
		visited[i][j] = true
		for _, dir := range dirs {
			next_i, next_j := i+dir[0], j+dir[1]
			dfseee(board, word, next_i, next_j, visited, pos+1)

		}
		visited[i][j] = false

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			found = false
			dfseee(board, word, i, j, visited, 0)
			if found {
				return true
			}
		}
	}
	return false
}

// @lc code=end

func TestWordSearch(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

*/
