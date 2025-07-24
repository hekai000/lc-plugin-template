/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=30104
 *
 * [416] 分割等和子集
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	// for j := 1; j <= target; j++ {
	// 	dp[0][j] = false
	// }
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][target]

}

// @lc code=end

func TestPartitionEqualSubsetSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/
