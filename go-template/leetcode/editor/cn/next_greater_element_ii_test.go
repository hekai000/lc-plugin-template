/*
 * @lc app=leetcode.cn id=503 lang=golang
 * @lcpr version=30104
 *
 * [503] 下一个更大元素 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	s := make([]int, 0)

	for i := 2*n - 1; i >= 0; i-- {
		for len(s) > 0 && nums[i%n] >= s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1]
		}
		s = append(s, nums[i%n])
	}
	return res
}

// @lc code=end

func TestNextGreaterElementIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,3]\n
// @lcpr case=end

*/
