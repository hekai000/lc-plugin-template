/*
 * @lc app=leetcode.cn id=438 lang=golang
 * @lcpr version=30103
 *
 * [438] 找到字符串中所有字母异位词
 */

package leetcode_solutions

import "testing"

// @lc code=start
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	for i := range p {
		need[p[i]]++
	}
	res := []int{}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
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
	return res
}

// @lc code=end

func TestFindAllAnagramsInAString(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/
