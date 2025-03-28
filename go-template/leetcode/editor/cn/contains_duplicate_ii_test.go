/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=30104
 *
 * [219] 存在重复元素 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, 0
	window := make(map[int]int)

	for right < len(nums) {
		c := nums[right]
		right++

		window[c]++

		if window[c] >= 2 {
			return true
		}

		for right-left > k {
			d := nums[left]
			left++

			window[d]--
		}

	}
	return false
}

// @lc code=end

func TestContainsDuplicateIi(t *testing.T) {
	// your test code here
	a := []int{1, 2, 3, 1}
	k := 3
	containsNearbyDuplicate(a, k)

}

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/
