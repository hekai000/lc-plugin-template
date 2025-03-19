/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=30103
 *
 * [14] 最长公共前缀
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestCommonPrefix(strs []string) string {
	m, n := len(strs), len(strs[0])

	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			if col == len(strs[row]) || strs[row][col] != strs[0][col] {
				return strs[0][:col]
			}
		}
	}
	return strs[0]
}

// @lc code=end

func TestLongestCommonPrefix(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/
