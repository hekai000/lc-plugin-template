/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=30104
 *
 * [322] 零钱兑换
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func coinChange2(coins []int, amount int) int {
	//dp[amount]返回目标金额为amount时的最少硬币数
	if amount == 0 {
		return 0
	}
	return dpcc(coins, amount)
}

func coinChange(coins []int, amount int) int {
	//dp[amount]返回目标金额为amount时的最少硬币数

	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -888
	}
	return dpcd(coins, amount, memo)
}

func dpcd(coins []int, amount int, memo []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	if memo[amount] != -888 {
		return memo[amount]
	}
	res := math.MaxInt32
	for _, coin := range coins {
		subProblem := dpcd(coins, amount-coin, memo)
		if subProblem == -1 {
			continue
		}
		res = mincc(res, 1+subProblem)
	}
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}
func dpcc(coins []int, amount int) int {
	dpa := make([]int, amount+1)
	for i := range dpa {
		dpa[i] = amount + 1
	}

	dpa[0] = 0
	for i := 0; i < len(dpa); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dpa[i] = mincc(dpa[i], 1+dpa[i-coin])
		}
	}
	if dpa[amount] == amount+1 {
		return -1
	}
	return dpa[amount]

}

func mincc(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func TestCoinChange(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
