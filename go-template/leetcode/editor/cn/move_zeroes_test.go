/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30100
 *
 * [283] 移动零
 */

package leetcode_solutions

import "testing"

// @lc code=start
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
	return
}

// @lc code=end

func TestMoveZeroes(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
