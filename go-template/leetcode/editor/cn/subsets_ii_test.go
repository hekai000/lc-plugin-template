/*
 * @lc app=leetcode.cn id=90 lang=golang
 * @lcpr version=30104
 *
 * [90] 子集 II
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	track := []int{}
	sort.Ints(nums)
	var backtrackswd func(start int)
	backtrackswd = func(start int) {
		result = append(result, append([]int(nil), track...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			backtrackswd(i + 1)
			track = track[:len(track)-1]

		}
	}
	backtrackswd(0)
	return result
}

// @lc code=end

func TestSubsetsIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
