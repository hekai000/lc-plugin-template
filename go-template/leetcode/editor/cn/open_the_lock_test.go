/*
 * @lc app=leetcode.cn id=752 lang=golang
 * @lcpr version=30104
 *
 * [752] 打开转盘锁
 */

package leetcode_solutions

import "testing"

// @lc code=start
func openLock(deadends []string, target string) int {
	start := "0000"
	visited := make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
		if deadend == "0000" {
			return -1
		}
	}

	queue := []string{start}
	step := 0
	visited[start] = true
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return step
			}
			for _, neighbor := range getNeighborsLock(cur) {
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

func getNeighborsLock(s string) []string {
	neighbors := make([]string, 0)
	for i := 0; i < 4; i++ {
		neighbors = append(neighbors, plusOne(s, i))
		neighbors = append(neighbors, minusOne(s, i))
	}
	return neighbors
}

func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}
func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j]--
	}
	return string(ch)
}

// @lc code=end

func TestOpenTheLock(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["0201","0101","0102","1212","2002"]\n"0202"\n
// @lcpr case=end

// @lcpr case=start
// ["8888"]\n"0009"\n
// @lcpr case=end

// @lcpr case=start
// ["8887","8889","8878","8898","8788","8988","7888","9888"]\n"8888"\n
// @lcpr case=end

*/
