/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30104
 *
 * [11] 盛最多水的容器
 */

package leetcode_solutions

import "testing"

// @lc code=start
func maxArea(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	left, right := 0, n-1
	for left < right {
		curArea := minaa(height[left], height[right]) * (right - left)
		res = maxaa(res, curArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func minaa(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxaa(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestContainerWithMostWater(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
