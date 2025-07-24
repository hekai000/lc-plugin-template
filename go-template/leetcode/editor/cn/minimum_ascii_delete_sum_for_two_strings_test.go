/*
 * @lc app=leetcode.cn id=712 lang=golang
 * @lcpr version=30104
 *
 * [712] 两个字符串的最小ASCII删除和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			if i == m {
				res := 0
				for z := j; z < n; z++ {
					res += int(s2[z])
				}
				memo[i][j] = res

			} else if j == n {
				res2 := 0
				for z := i; z < m; z++ {
					res2 += int(s1[z])
				}
				memo[i][j] = res2
			} else {
				memo[i][j] = -1
			}

		}
	}
	return dpmmd(s1, 0, s2, 0, memo)
}

func dpmmd(s1 string, i int, s2 string, j int, memo [][]int) int {
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if s1[i] == s2[j] {
		memo[i][j] = dpmmd(s1, i+1, s2, j+1, memo)
	} else {
		memo[i][j] = minmmd(int(s1[i])+dpmmd(s1, i+1, s2, j, memo), int(s2[j])+dpmmd(s1, i, s2, j+1, memo))
	}
	return memo[i][j]
}

func minmmd(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func TestMinimumAsciiDeleteSumForTwoStrings(t *testing.T) {
	// your test code here
	s1 := "sea"
	s2 := "eat"
	minimumDeleteSum(s1, s2)

}

/*
// @lcpr case=start
// "sea"\n"eat"\n
// @lcpr case=end

// @lcpr case=start
// "delete"\n"leet"\n
// @lcpr case=end

*/
