/*
 * @lc app=leetcode.cn id=410 lang=golang
 * @lcpr version=30104
 *
 * [410] 分割数组的最大值
 */

package leetcode_solutions

import "testing"

// @lc code=start
func splitArray(nums []int, k int) int {
	//假设和的最大值至少为x,那么f(x)表示可以分成f(x)堆
	left, right := 0, 1
	for _, w := range nums {
		if left < w {
			left = w
		}
		right += w
	}

	for left <= right {
		mid := left + (right-left)/2
		if fkk(nums, mid) <= k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func fkk(nums []int, s int) int {
	d := 0
	for i := 0; i < len(nums); {
		cap := s
		for i < len(nums) {
			if cap < nums[i] {
				break
			} else {
				cap -= nums[i]
			}
			i++
		}
		d++
	}
	return d
}

// @lc code=end

func TestSplitArrayLargestSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [7,2,5,10,8]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,4]\n3\n
// @lcpr case=end

*/
