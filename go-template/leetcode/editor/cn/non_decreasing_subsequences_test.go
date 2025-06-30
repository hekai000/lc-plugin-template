/*
 * @lc app=leetcode.cn id=491 lang=golang
 * @lcpr version=30104
 *
 * [491] 非递减子序列
 */

package leetcode_solutions

import "testing"

// @lc code=start
func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	var backtrackfss func(start int)
	backtrackfss = func(start int) {

		if len(track) >= 2 {
			result = append(result, append([]int(nil), track...))
		}
		//used map只负责本层元素是否被使用
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {

			if (len(track) != 0) && (nums[i] < track[len(track)-1]) {
				continue
			}
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			track = append(track, nums[i])
			backtrackfss(i + 1)
			track = track[:len(track)-1]
			//used[nums[i]] = false

		}
	}
	backtrackfss(0)
	return result
}

// @lc code=end

func TestNonDecreasingSubsequences(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,6,7,7]\n
// @lcpr case=end

// @lcpr case=start
// [4,4,3,2,1]\n
// @lcpr case=end

*/
