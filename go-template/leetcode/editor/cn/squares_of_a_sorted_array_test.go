/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30100
 *
 * [977] 有序数组的平方
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func sortedSquares2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums

}

func sortedSquares(nums []int) []int {

	n := len(nums)
	i, j := 0, n-1
	p := n - 1
	res := make([]int, n)
	for i <= j {
		if abs(nums[i]) > abs(nums[j]) {
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
		p--
	}
	return res
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

// @lc code=end

func TestSquaresOfASortedArray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end

// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/
