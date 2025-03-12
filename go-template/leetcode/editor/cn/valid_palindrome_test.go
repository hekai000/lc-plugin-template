/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30100
 *
 * [125] 验证回文串
 */

package leetcode_solutions

import (
	"strings"
	"testing"
	"unicode"
)

// @lc code=start
func isPalindrome(s string) bool {
	output := cleanString(s)

	left, right := 0, len(output)-1

	for left < right {
		if output[left] != output[right] {
			return false
		}
		left++
		right--
	}
	return true

}
func cleanString(s string) string {
	var builder strings.Builder
	builder.Grow(len(s)) // 预分配空间以提高性能

	for _, r := range s {
		// 将字符转换为小写
		lowerR := unicode.ToLower(r)
		// 检查是否为小写字母（a-z）或原字符是数字（0-9）
		if (lowerR >= 'a' && lowerR <= 'z') || (r >= '0' && r <= '9') {
			builder.WriteRune(lowerR)
		}
	}

	return builder.String()
}

// @lc code=end

func TestValidPalindrome(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/
