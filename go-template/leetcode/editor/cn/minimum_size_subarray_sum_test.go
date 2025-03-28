/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30104
 *
 * [209] 长度最小的子数组
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	windowSum := 0
	minLength := math.MaxInt32

	for right < len(nums) {
		c := nums[right]
		right++

		windowSum += c

		for left < right && windowSum >= target {
			if right-left < minLength {
				minLength = right - left
			}
			d := nums[left]
			left++

			windowSum -= d

		}

	}
	if minLength == math.MaxInt32 {
		minLength = 0
	}
	return minLength
}

// @lc code=end

func TestMinimumSizeSubarraySum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/
