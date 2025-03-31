/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=30104
 *
 * [704] 二分查找
 */

package leetcode_solutions

import "testing"

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end

func TestBinarySearch(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

*/
