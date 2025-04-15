/*
 * @lc app=leetcode.cn id=402 lang=golang
 * @lcpr version=30104
 *
 * [402] 移掉 K 位数字
 */

package leetcode_solutions

import "testing"

// @lc code=start
func removeKdigits(num string, k int) string {
	stk := []rune{}
	for _, c := range num {
		for len(stk) > 0 && c < stk[len(stk)-1] && k > 0 {
			k--
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 && c == '0' {
			continue
		}
		stk = append(stk, c)
	}
	for k > 0 && len(stk) > 0 {
		stk = stk[:len(stk)-1]
		k--
	}
	if len(stk) == 0 {
		return "0"
	}
	return string(stk)
}

// @lc code=end

func TestRemoveKDigits(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "1432219"\n3\n
// @lcpr case=end

// @lcpr case=start
// "10200"\n1\n
// @lcpr case=end

// @lcpr case=start
// "10"\n2\n
// @lcpr case=end

*/
