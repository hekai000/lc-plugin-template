/*
 * @lc app=leetcode.cn id=150 lang=golang
 * @lcpr version=30104
 *
 * [150] 逆波兰表达式求值
 */

package leetcode_solutions

import (
	"strconv"
	"testing"
)

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		if isInteger(token) {
			tokenI, _ := strconv.Atoi(token)
			stack = append(stack, tokenI)
		} else {
			a1, a2 := stack[len(stack)-1], stack[len(stack)-2]

			stack = stack[:len(stack)-2]
			res := 0
			if token == "*" {
				res = a1 * a2
			} else if token == "/" {
				res = a2 / a1
			} else if token == "+" {
				res = a1 + a2
			} else if token == "-" {
				res = a2 - a1
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
func isInteger(s string) bool {
	_, err := strconv.Atoi(s) // 或 strconv.ParseInt(s, 10, 64)
	return err == nil
}

// @lc code=end

func TestEvaluateReversePolishNotation(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["2","1","+","3","*"]\n
// @lcpr case=end

// @lcpr case=start
// ["4","13","5","/","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]\n
// @lcpr case=end

*/
