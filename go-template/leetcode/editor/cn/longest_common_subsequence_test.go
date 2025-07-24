/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=30104
 *
 * [1143] 最长公共子序列
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			if i == m {
				memo[i][j] = 0
			} else if j == n {
				memo[i][j] = 0
			} else {
				memo[i][j] = -1
			}

		}
	}

	res := dplcs(text1, 0, text2, 0, memo)
	return res
}

func dplcs(text1 string, i int, text2 string, j int, memo [][]int) int {

	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if text1[i] == text2[j] {
		memo[i][j] = dplcs(text1, i+1, text2, j+1, memo) + 1
	} else {
		memo[i][j] = maxlcs(dplcs(text1, i+1, text2, j, memo), dplcs(text1, i, text2, j+1, memo))
	}
	return memo[i][j]
}

func maxlcs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestLongestCommonSubsequence(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/
