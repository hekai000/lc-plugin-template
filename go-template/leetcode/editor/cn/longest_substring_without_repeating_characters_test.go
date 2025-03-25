/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30103
 *
 * [3] 无重复字符的最长子串
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++

		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		if right-left > res {
			res = right - left
		}
	}
	return res

}
func lengthOfLongestSubstring1(s string) int {
	window := make(map[byte]int)
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	left, right := 0, 0

	length := math.MinInt32

	for right < len(s) {
		c := s[right]
		right++

		window[c]++

		if checkOne(window) == true {
			if len(window) > length {
				length = len(window)
			}

		}
		for checkOne(window) == false {
			if right-left > length {
				length = right - left - 1
			}

			d := s[left]
			left++
			window[d]--
			if window[d] == 0 {
				delete(window, d)
			}

		}

	}
	return length
}

func checkOne(m map[byte]int) bool {
	for _, v := range m {
		if v >= 2 {
			return false
		}
	}
	return true
}

// @lc code=end

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// your test code here
	s := "nfpdmpi"

	lengthOfLongestSubstring(s)
}

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
