/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=30100
 *
 * [80] 删除有序数组中的重复项 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
func removeDuplicates(nums []int) int {
	imap := make(map[int]int, 0)

	slow, fast := 0, 0
	for fast < len(nums) {
		if imap[nums[fast]] < 2 {
			imap[nums[fast]] += 1
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end

func TestRemoveDuplicatesFromSortedArrayIi(t *testing.T) {
	// your test code here
	arr := []int{0, 0, 1, 1, 1, 1, 2}
	removeDuplicates(arr)

}

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/
