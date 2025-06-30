/*
 * @lc app=leetcode.cn id=967 lang=golang
 * @lcpr version=30104
 *
 * [967] 连续差相同的数字
 */

package leetcode_solutions

import (
	"strconv"
	"testing"
)

// @lc code=start
// numsSameConsecDiff 函数用于返回所有长度为n且相邻数字的差为k的数字
func numsSameConsecDiff2(n int, k int) []int {
	result := []int{}
	track := []rune{}

	var backtracknscd func(index int)
	backtracknscd = func(index int) {
		if index == n {
			temp, _ := strconv.Atoi(string(track))

			result = append(result, temp)
			return
		}

		for ch := '0'; ch <= '9'; ch++ {
			if (index == n-1) && ch == '0' {
				continue
			}
			if !isValidSC(track, ch, k) {
				continue
			}
			track = append([]rune{ch}, track...)
			backtracknscd(index + 1)
			track = track[1:]
		}

	}
	backtracknscd(0)
	return result

}

func numsSameConsecDiff(n int, k int) []int {
	result := []int{}
	track := 0
	digit := 0
	backtracknscd2(n, k, &result, &track, &digit)
	return result
}

func backtracknscd2(n int, k int, result *[]int, track *int, digit *int) {
	if n == *digit {
		*result = append(*result, *track)
		return
	}
	for i := 0; i <= 9; i++ {
		if *digit == 0 && i == 0 {
			continue
		}

		if *digit > 0 && absnscd(i-(*track%10)) != k {
			continue
		}
		*digit++
		*track = *track*10 + i
		backtracknscd2(n, k, result, track, digit)
		*track /= 10
		*digit--
	}
	
}

func absnscd(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func isValidSC(track []rune, ch rune, k int) bool {
	if len(track) == 0 {
		return true
	}
	gap := int(ch - track[0])
	if gap < 0 {
		gap = -gap
	}
	if gap == k {
		return true
	}
	return false
}

// @lc code=end

func TestNumbersWithSameConsecutiveDifferences(t *testing.T) {
	// your test code here

	numsSameConsecDiff(3, 1)
}

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 2\n1\n
// @lcpr case=end

// @lcpr case=start
// 2\n0\n
// @lcpr case=end

// @lcpr case=start
// 2\n2\n
// @lcpr case=end

*/
