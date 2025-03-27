/*
 * @lc app=leetcode.cn id=424 lang=golang
 * @lcpr version=30104
 *
 * [424] 替换后的最长重复字符
 */

package leetcode_solutions

import "testing"

// @lc code=start

func characterReplacement(s string, k int) int {
	left, right := 0, 0
	window := make([]int, 26)
	windowMaxCount := 0
	res := 0
	for right < len(s) {
		c := s[right] - 'A'
		right++

		window[c]++
		if window[c] > windowMaxCount {
			windowMaxCount = window[c]
		}

		for left < right && right-left-windowMaxCount > k {
			d := s[left] - 'A'
			left++

			window[d]--

		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}
func characterReplacement1(s string, k int) int {
	left, right := 0, 0
	window := make(map[byte]int)

	length := len(s)
	res := 0
	for right < length {
		c := s[right]
		right++

		window[c]++

		for left < right && right-left-mostC(window) > k {
			d := s[left]
			left++

			window[d]--
		}
		res = maxC(res, mysumM(window))
	}
	return res
}

func mysumM(m map[byte]int) int {
	sum := 0
	for _, j := range m {
		sum += j
	}
	return sum
}

func maxC(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func mostC(m map[byte]int) int {
	res := 0

	for _, j := range m {
		if j > res {
			res = j
		}
	}
	return res
}

// @lc code=end

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	// your test code here
	s := "ABAB"
	k := 2
	characterReplacement(s, k)

}

/*
// @lcpr case=start
// "ABAB"\n2\n
// @lcpr case=end

// @lcpr case=start
// "AABABBA"\n1\n
// @lcpr case=end

*/
