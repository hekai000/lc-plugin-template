/*
 * @lc app=leetcode.cn id=292 lang=golang
 * @lcpr version=30104
 *
 * [292] Nim 游戏
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canWinNim(n int) bool {
	return n%4 != 0
}

// @lc code=end

func TestNimGame(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
