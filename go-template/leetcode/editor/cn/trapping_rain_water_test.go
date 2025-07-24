/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30104
 *
 * [42] 接雨水
 */

package leetcode_solutions

import "testing"

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	res := 0
	l_max := make([]int, n)
	r_max := make([]int, n)
	l_max[0] = height[0]
	r_max[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		l_max[i] = maxtrap(l_max[i-1], height[i])
	}
	for j := n - 2; j >= 0; j-- {
		r_max[j] = maxtrap(r_max[j+1], height[j])
	}

	for i := 0; i < n; i++ {
		res += mintrap(l_max[i], r_max[i]) - height[i]
	}
	return res
}

func maxtrap(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mintrap(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func TestTrappingRainWater(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
