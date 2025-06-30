/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30104
 *
 * [39] 组合总和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	track := []int{}
	trackSum := 0

	var backtrackcs func(start int)

	backtrackcs = func(start int) {
		if trackSum > target {
			return
		}
		if trackSum == target {
			result = append(result, append([]int(nil), track...))
			return
		}

		for i := start; i < len(candidates); i++ {
			trackSum += candidates[i]
			track = append(track, candidates[i])

			backtrackcs(i)

			trackSum -= candidates[i]
			track = track[:len(track)-1]
		}
	}
	backtrackcs(0)
	return result
}

// @lc code=end

func TestCombinationSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
