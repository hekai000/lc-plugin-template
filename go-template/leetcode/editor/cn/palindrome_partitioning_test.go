/*
 * @lc app=leetcode.cn id=131 lang=golang
 * @lcpr version=30104
 *
 * [131] 分割回文串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func partition(s string) [][]string {
	result := [][]string{}
	track := []string{}

	backtrackpt(s, 0, &result, &track)
	return result
}

func backtrackpt(s string, start int, result *[][]string, track *[]string) {
	if start == len(s) {
		*result = append(*result, append([]string(nil), *track...))
		return
	}
	for i := start; i < len(s); i++ {
		if !isPalindstring(s, start, i) {
			continue
		}
		*track = append(*track, s[start:i+1])
		backtrackpt(s, i+1, result, track)
		*track = (*track)[:len(*track)-1]
	}
}

func isPalindstring(s string, start int, i int) bool {
	for start < i {
		if s[start] != s[i] {
			return false
		}
		start++
		i--
	}
	return true
}

// @lc code=end

func TestPalindromePartitioning(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "aab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/
