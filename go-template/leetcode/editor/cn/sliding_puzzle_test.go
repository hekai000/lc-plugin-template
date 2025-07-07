/*
 * @lc app=leetcode.cn id=773 lang=golang
 * @lcpr version=30104
 *
 * [773] 滑动谜题
 */

package leetcode_solutions

import (
	"strings"
	"testing"
)

// @lc code=start
func slidingPuzzle(board [][]int) int {
	start := ""
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			start += string(rune(board[i][j]) + '0')
		}
	}
	target := "123450"

	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true
	step := 0

	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			for _, neighbor := range getNeighbors(cur) {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
		step++
	}
	return -1

}

func getNeighbors(board string) []string {
	mapping := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}
	idx := strings.Index(board, "0")
	neighors := []string{}
	for _, adj := range mapping[idx] {
		newBoard := swapBoard(board, idx, adj)
		neighors = append(neighors, newBoard)
	}
	return neighors
}

func swapBoard(board string, i, j int) string {
	chars := []rune(board)
	chars[i], chars[j] = chars[j], chars[i]
	return string(chars)
}

// @lc code=end

func TestSlidingPuzzle(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,2,3],[4,0,5]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[5,4,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[4,1,2],[5,0,3]]\n
// @lcpr case=end

*/
