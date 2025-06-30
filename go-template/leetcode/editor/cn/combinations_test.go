/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30104
 *
 * [77] 组合
 */

package leetcode_solutions

import "testing"

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}
	track := []int{}
	var backtracecb func(start int)
	backtracecb = func(start int) {
		if k == len(track) {
			result = append(result, append([]int(nil), track...))
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backtracecb(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtracecb(1)
	return result
}

// @lc code=end

func TestCombinations(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
