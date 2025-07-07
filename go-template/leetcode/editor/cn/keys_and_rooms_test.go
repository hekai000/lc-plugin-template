/*
 * @lc app=leetcode.cn id=841 lang=golang
 * @lcpr version=30104
 *
 * [841] 钥匙和房间
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canVisitAllRooms2(rooms [][]int) bool {
	n := len(rooms)
	visited := make(map[int]bool, n)
	visitedCount := 0
	queue := []int{0}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if visited[cur] == true {
			continue
		}
		visited[cur] = true
		visitedCount++
		for _, v := range rooms[cur] {
			if visited[v] == true {
				continue
			} else {
				queue = append(queue, v)
			}

		}

	}
	if visitedCount == n {
		return true
	}
	return false

}
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nextRoom := range rooms[cur] {
			if !visited[nextRoom] {
				visited[nextRoom] = true
				queue = append(queue, nextRoom)
			}
		}
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true

}

// @lc code=end

func TestKeysAndRooms(t *testing.T) {
	// your test code here
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	canVisitAllRooms(rooms)
}

/*
// @lcpr case=start
// [[1],[2],[3],[]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[3,0,1],[2],[0]]\n
// @lcpr case=end

*/
