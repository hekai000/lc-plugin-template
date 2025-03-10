/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=30100
 *
 * [26] 删除有序数组中的重复项
 */

package leetcode_solutions

import "testing"

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	k := len(nums)
	if k < 2 {
		return k
	}

	for fast < k {
		if nums[slow] != nums[fast] {

			slow++
			nums[slow] = nums[fast]

		}
		fast++
	}
	return slow + 1
}

// @lc code=end

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,2,2,3,3,4]\n
// @lcpr case=end

*/
