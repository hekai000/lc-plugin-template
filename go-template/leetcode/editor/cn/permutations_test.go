/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=30104
 *
 * [46] 全排列
 */

package leetcode_solutions

import "testing"

// @lc code=start
func permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	backtrackp(nums, path, used, &result)
	return result
}

func backtrackp(nums []int, path []int, used []bool, result *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrackp(nums, path, used, result)

		path = path[:len(path)-1]
		used[i] = false

	}

}

// @lc code=end

func TestPermutations(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
