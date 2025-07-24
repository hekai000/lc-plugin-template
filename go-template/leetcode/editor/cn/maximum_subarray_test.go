/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30104
 *
 * [53] 最大子数组和
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func maxSubArray(nums []int) int {
	m := len(nums)
	preSum := make([]int, m+1)
	preSum[0] = 0
	for i := 1; i <= m; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	res := math.MinInt32
	minVal := math.MaxInt32
	for i := 0; i < m; i++ {
		minVal = minmsa(minVal, preSum[i])
		res = maxmsa(res, preSum[i+1]-minVal)
	}
	return res

}
func maxSubArray4(nums []int) int {
	m := len(nums)
	left, right := 0, 0
	windowSum := 0
	maxS := math.MinInt32
	for right < m {
		windowSum += nums[right]
		right++

		if windowSum > maxS {
			maxS = windowSum
		}
		for windowSum < 0 {
			windowSum -= nums[left]
			left++
		}
	}
	return maxS
}

func maxSubArray3(nums []int) int {
	m := len(nums)
	res := math.MinInt32
	dp := make([]int, m)
	if m == 0 {
		return 0
	}
	dp[0] = nums[0]
	for i := 1; i < m; i++ {
		dp[i] = maxmsa(nums[i], nums[i]+dp[i-1])
	}
	for i := 0; i < m; i++ {
		res = maxmsa(res, dp[i])
	}
	return res
}

func maxSubArray2(nums []int) int {
	m := len(nums)
	maxS := math.MinInt32
	memo := make([]int, m)
	for i := range memo {
		memo[i] = -10001
	}

	for i := 0; i < m; i++ {
		maxS = maxmsa(maxS, dpmsa(nums, i, &memo))
	}
	return maxS
}

func dpmsa(nums []int, i int, memo *[]int) int {

	if i == 0 {
		return nums[0]
	}
	var sub1 int
	if (*memo)[i] != -10001 {
		sub1 = (*memo)[i]
	} else {
		sub1 = dpmsa(nums, i-1, memo) + nums[i]
	}
	(*memo)[i] = maxmsa(nums[i], sub1)
	return (*memo)[i]
}

func maxmsa(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minmsa(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func TestMaximumSubarray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
