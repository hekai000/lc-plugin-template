/*
 * @lc app=leetcode.cn id=365 lang=golang
 * @lcpr version=30104
 *
 * [365] 水壶问题
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canMeasureWater(x int, y int, target int) bool {
	queue := [][]int{}
	visited := make(map[int64]bool)

	queue = append(queue, []int{0, 0})
	visited[int64(0)*(int64(y)+1)+0] = true

	for len(queue) > 0 {
		curState := queue[0]
		queue = queue[1:]
		if curState[0] == target || curState[1] == target || curState[0]+curState[1] == target {
			return true
		}
		nextStates := [][]int{
			// 把 1 桶灌满
			{x, curState[1]},
			// 把 2 桶灌满
			{curState[0], y},
			// 把 1 桶倒空
			{0, curState[1]},
			// 把 2 桶倒空
			{curState[0], 0},
			// 把 1 桶的水灌进 2 桶，直到 1 桶空了或者 2 桶满了
			{curState[0] - min(curState[0], y-curState[1]), curState[1] + min(curState[0], y-curState[1])},
			// 把 2 桶的水灌进 1 桶，直到 2 桶空了或者 1 桶满了
			{curState[0] + min(curState[1], x-curState[0]), curState[1] - min(curState[1], x-curState[0])},
		}
		for _, nextState := range nextStates {
			hash := int64(nextState[0])*(int64(y)+1) + int64(nextState[1])
			if visited[hash] {
				continue
			}
			queue = append(queue, nextState)
			visited[hash] = true
		}
	}
	return false

}

// @lc code=end

func TestWaterAndJugProblem(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 3\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// 2\n6\n5\n
// @lcpr case=end

// @lcpr case=start
// 1\n2\n3\n
// @lcpr case=end

*/
