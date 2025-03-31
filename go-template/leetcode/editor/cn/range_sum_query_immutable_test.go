/*
 * @lc app=leetcode.cn id=303 lang=golang
 * @lcpr version=30104
 *
 * [303] 区域和检索 - 数组不可变
 */

package leetcode_solutions

import "testing"

// @lc code=start
type NumArray struct {
	PreSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	return NumArray{PreSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {

	return this.PreSum[right+1] - this.PreSum[left]
}

// type NumArray2 struct {
// 	numsA []int
// }

// func Constructor2(nums []int) NumArray {
// 	a := NumArray{numsA: nums}
// 	return a
// }

// func (this *NumArray) SumRange2(left int, right int) int {
// 	sum := 0
// 	for i := left; i <= right; i++ {
// 		sum += this.numsA[i]
// 	}
// 	return sum
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

func TestRangeSumQueryImmutable(t *testing.T) {
	// your test code here
	b := []int{-2, 0, 3, -5, 2, -1}
	c := Constructor(b)
	c.SumRange(0, 2)

}

/*
// @lcpr case=start
// ["NumArray", "SumRange", "SumRange", "SumRange"]\n[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]\n
// @lcpr case=end

*/
