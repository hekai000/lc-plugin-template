/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=30103
 *
 * [76] 最小覆盖子串
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func minWindow(s string, t string) string {

	left, right := 0, 0
	window := make(map[byte]int)
	need := make(map[byte]int)
	valid := 0
	start, length := 0, math.MaxInt32
	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {

			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {

				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
func minWindow1(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)
	valid := 0
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}

		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]

}

// @lc code=end

func TestMinimumWindowSubstring(t *testing.T) {
	// your test code here
	sg := "aaa"
	tg := "aa"

	minWindow(sg, tg)

}

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
