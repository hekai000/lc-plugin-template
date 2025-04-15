/*
 * @lc app=leetcode.cn id=581 lang=golang
 * @lcpr version=30104
 *
 * [581] 最短无序连续子数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := n, -1
	incrStk := make([]int, 0)
	decrStk := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(incrStk) > 0 && nums[i] < nums[incrStk[len(incrStk)-1]] {
			left = min(left, incrStk[len(incrStk)-1])
			incrStk = incrStk[:len(incrStk)-1]
		}
		incrStk = append(incrStk, i)
	}

	for i := n - 1; i >= 0; i-- {
		for len(decrStk) > 0 && nums[i] > nums[decrStk[len(decrStk)-1]] {
			right = max(right, decrStk[len(decrStk)-1])
			decrStk = decrStk[:len(decrStk)-1]
		}
		decrStk = append(decrStk, i)
	}
	if left == n && right == -1 {
		return 0
	}
	return right - left + 1
}

// @lc code=end

func TestShortestUnsortedContinuousSubarray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,6,4,8,10,9,15]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
