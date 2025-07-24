/*
 * @lc app=leetcode.cn id=319 lang=golang
 * @lcpr version=30104
 *
 * [319] 灯泡开关
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

// @lc code=end

func TestBulbSwitcher(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
