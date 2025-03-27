/*
 * @lc app=leetcode.cn id=1004 lang=golang
 * @lcpr version=30103
 *
 * [1004] 最大连续1的个数 III
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestOnes1(nums []int, k int) int {

	left, right := 0, 0

	oneCounts := 0
	res := 0
	for right < len(nums) {
		if nums[right] == 1 {
			oneCounts++
		}
		right++

		for right-left-oneCounts > k {
			if nums[left] == 1 {
				oneCounts--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	zeros := 0
	res := 0

	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}
		right++
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestMaxConsecutiveOnesIii(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,1,1,0,0,0,1,1,1,1,0]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]\n3\n
// @lcpr case=end

*/
