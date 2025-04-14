/*
 * @lc app=leetcode.cn id=739 lang=golang
 * @lcpr version=30104
 *
 * [739] 每日温度
 */

package leetcode_solutions

import "testing"

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	s := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[i] >= temperatures[s[len(s)-1]] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}

// @lc code=end

func TestDailyTemperatures(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [73,74,75,71,69,72,76,73]\n
// @lcpr case=end

// @lcpr case=start
// [30,40,50,60]\n
// @lcpr case=end

// @lcpr case=start
// [30,60,90]\n
// @lcpr case=end

*/
