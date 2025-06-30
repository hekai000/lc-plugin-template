/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=30104
 *
 * [47] 全排列 II
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	used := make(map[int]bool, len(nums))
	sort.Ints(nums)

	var backtrackpu func()
	backtrackpu = func() {
		if len(track) == len(nums) {
			result = append(result, append([]int(nil), track...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			track = append(track, nums[i])
			used[i] = true
			backtrackpu()
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrackpu()
	return result
}

// @lc code=end

func TestPermutationsIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
