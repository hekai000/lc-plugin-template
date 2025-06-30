/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=30104
 *
 * [78] 子集
 */

package leetcode_solutions

import "testing"

// @lc code=start
func subsets(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	var backtracess func(start int)
	backtracess = func(start int) {
		result = append(result, append([]int(nil), track...))
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtracess(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtracess(0)
	return result
}

// @lc code=end

func TestSubsets(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
