/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30104
 *
 * [300] 最长递增子序列
 */

package leetcode_solutions

import "testing"

// @lc code=start
func lengthOfLIS(nums []int) int {
	dpl := make([]int, len(nums))
	for i := range dpl {
		dpl[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dpl[i] = maxl(dpl[i], dpl[j]+1)
			}
		}
	}
	res := 0
	for _, v := range dpl {
		res = maxl(res, v)
	}
	return res
}

func maxl(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestLongestIncreasingSubsequence(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/
