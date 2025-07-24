/*
 * @lc app=leetcode.cn id=583 lang=golang
 * @lcpr version=30104
 *
 * [583] 两个字符串的删除操作
 */

package leetcode_solutions

import "testing"

// @lc code=start
func minDistance22(word1 string, word2 string) int {
	lcs := longestCommonSubsequence2(word1, word2)
	m, n := len(word1), len(word2)
	return m + n - 2*lcs
}

func longestCommonSubsequence2(text1 string, text2 string) int {
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

	res := dplcs2(text1, 0, text2, 0, memo)
	return res
}

func dplcs2(text1 string, i int, text2 string, j int, memo [][]int) int {

	if memo[i][j] != -1 {
		return memo[i][j]
	}
	if text1[i] == text2[j] {
		memo[i][j] = dplcs2(text1, i+1, text2, j+1, memo) + 1
	} else {
		memo[i][j] = maxlcs2(dplcs2(text1, i+1, text2, j, memo), dplcs2(text1, i, text2, j+1, memo))
	}
	return memo[i][j]
}

func maxlcs2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestDeleteOperationForTwoStrings(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "sea"\n"eat"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"etco"\n
// @lcpr case=end

*/
