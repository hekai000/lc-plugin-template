/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=30100
 *
 * [5] 最长回文子串
 */

package leetcode_solutions

import "testing"

// @lc code=start

func findPalindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		temp1 := findPalindrome(s, i, i)
		temp2 := findPalindrome(s, i, i+1)
		if len(temp1) > len(res) {
			res = temp1
		}
		if len(temp2) > len(res) {
			res = temp2
		}
	}
	return res
}

// @lc code=end

func TestLongestPalindromicSubstring(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/
