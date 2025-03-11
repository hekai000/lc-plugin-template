/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=30100
 *
 * [344] 反转字符串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return
}

// @lc code=end

func TestReverseString(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/
