/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=30104
 *
 * [72] 编辑距离
 */

package leetcode_solutions

import "testing"

// @lc code=start

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3dp(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dpmd(word1, m-1, word2, n-1, memo)
}

func dpmd(word1 string, i int, word2 string, j int, memo [][]int) int {
	if i == -1 {
		return j + 1
	}

	if j == -1 {
		return i + 1
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if word1[i] == word2[j] {
		memo[i][j] = dpmd(word1, i-1, word2, j-1, memo)
	} else {
		memo[i][j] = min3dp(dpmd(word1, i, word2, j-1, memo)+1,
			dpmd(word1, i-1, word2, j, memo)+1,
			dpmd(word1, i-1, word2, j-1, memo)+1)
	}
	return memo[i][j]
}

func min3dp(a, b, c int) int {
	if a > b {
		if b < c {
			return b
		} else {
			return c
		}
	} else {
		if a < c {
			return a
		} else {
			return c
		}
	}
}

// @lc code=end

func TestEditDistance(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "horse"\n"ros"\n
// @lcpr case=end

// @lcpr case=start
// "intention"\n"execution"\n
// @lcpr case=end

*/
