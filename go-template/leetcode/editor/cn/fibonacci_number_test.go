/*
 * @lc app=leetcode.cn id=509 lang=golang
 * @lcpr version=30104
 *
 * [509] 斐波那契数
 */

package leetcode_solutions

import "testing"

// @lc code=start
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return dp(n)
}

func dp(n int) int {
	dpa := make([]int, n+1)
	dpa[0] = 0
	dpa[1] = 1

	for i := 2; i <= n; i++ {
		dpa[i] = dpa[i-1] + dpa[i-2]
	}
	return dpa[n]
}

// @lc code=end

func TestFibonacciNumber(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/
