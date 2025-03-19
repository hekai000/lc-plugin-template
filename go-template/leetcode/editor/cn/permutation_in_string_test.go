/*
 * @lc app=leetcode.cn id=567 lang=golang
 * @lcpr version=30103
 *
 * [567] 字符串的排列
 */

package leetcode_solutions

import "testing"

// @lc code=start

func sum(m map[byte]int) int {
	res := 0
	for _, val := range m {
		res += val
	}
	return res
}
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[rune]int)
	window := make(map[rune]int)
	left, right := 0, 0
	valid := 0

	for _, c := range s1 {
		need[c]++
	}
	for right < len(s2) {
		c := rune(s2[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := rune(s2[left])
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--

			}

		}
	}
	return false
}
func checkInclusion1(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0

	for i := range s1 {
		need[s1[i]]++
	}
	for right < len(s2) {
		c := s2[right]
		right++

		window[c]++
		if window[c] == need[c] {
			valid++
		}

		for valid == len(need) {
			if len(s1) == sum(window) {
				return true
			}
			d := s2[left]
			left++

			if window[d] == need[d] {
				valid--
			}
			window[d]--

		}
	}
	return false
}

// @lc code=end

func TestPermutationInString(t *testing.T) {
	// your test code here
	s1 := "ba"
	s2 := "eidboaoo"
	checkInclusion(s1, s2)
}

/*
// @lcpr case=start
// "eidbaooo"\n
// @lcpr case=end

// @lcpr case=start
// "eidboaoo"\n
// @lcpr case=end

*/
