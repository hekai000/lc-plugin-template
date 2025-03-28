/*
 * @lc app=leetcode.cn id=395 lang=golang
 * @lcpr version=30104
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestSubstring(s string, k int) int {
	var len int
	for i := 1; i <= 26; i++ {
		len = max(len, longestKLetterSubstr(s, k, i))
	}
	return len
}

func longestKLetterSubstr(s string, k int, count int) int {
	left, right := 0, 0

	windowCount := make([]int, 26)
	windowValidCount := 0
	windowUniqCount := 0
	res := 0
	for right < len(s) {
		c := s[right]
		if windowCount[c-'a'] == 0 {
			windowUniqCount++
		}
		windowCount[c-'a']++
		if windowCount[c-'a'] == k {
			windowValidCount++
		}
		right++

		for windowUniqCount > count {
			d := s[left]
			if windowCount[d-'a'] == k {
				windowValidCount--
			}
			windowCount[d-'a']--
			if windowCount[d-'a'] == 0 {
				windowUniqCount--
			}
			left++
		}
		if windowValidCount == count {
			if right-left > res {
				res = right - left
			}
		}
	}
	return res
}

// @lc code=end

func TestLongestSubstringWithAtLeastKRepeatingCharacters(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "aaabb"\n3\n
// @lcpr case=end

// @lcpr case=start
// "ababbc"\n2\n
// @lcpr case=end

*/
