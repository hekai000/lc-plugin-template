/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=30104
 *
 * [55] 跳跃游戏
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		farthest = maxcj(farthest, i+nums[i])
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}

func maxcj(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestJumpGame(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
