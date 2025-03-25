/*
 * @lc app=leetcode.cn id=713 lang=golang
 * @lcpr version=30103
 *
 * [713] 乘积小于 K 的子数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	n := len(nums)
	res := [][]int{}
	windowP := 1
	count := 0
	for right < n {

		windowP *= nums[right]
		right++

		for left < right && windowP >= k {

			windowP /= nums[left]
			left++
		}
		count += right - left
		res = append(res, []int{left, right})

	}
	return count
}

// @lc code=end

func TestSubarrayProductLessThanK(t *testing.T) {
	// your test code here
	a := []int{10, 5, 2, 6}

	numSubarrayProductLessThanK(a, 100)

}

/*
// @lcpr case=start
// [10,5,2,6]\n100\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n0\n
// @lcpr case=end

*/
