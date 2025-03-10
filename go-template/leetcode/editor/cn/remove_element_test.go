/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30100
 *
 * [27] 移除元素
 */

package leetcode_solutions

import "testing"

// @lc code=start
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {

			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end

func TestRemoveElement(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/
