/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30104
 *
 * [20] 有效的括号
 */

package leetcode_solutions

import "testing"

// @lc code=start
func isValid(s string) bool {
	stack := []string{}
	for _, v := range s {
		if string(v) == "(" || string(v) == "[" || string(v) == "{" {
			stack = append(stack, string(v))
		} else {
			if len(stack) == 0 {
				return false
			}
			if string(v) == ")" {
				if len(stack) > 0 && stack[len(stack)-1] != "(" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			if string(v) == "]" {
				if len(stack) > 0 && stack[len(stack)-1] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
			if string(v) == "}" {
				if len(stack) > 0 && stack[len(stack)-1] != "{" {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}

	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end

func TestValidParentheses(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

*/
