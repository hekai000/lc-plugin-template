/*
 * @lc app=leetcode.cn id=331 lang=golang
 * @lcpr version=30104
 *
 * [331] 验证二叉树的前序序列化
 */

package leetcode_solutions

import (
	"strings"
	"testing"
)

// @lc code=start
func isValidSerialization(preorder string) bool {

	nodes := strings.Split(preorder, ",")
	// if nodes[len(nodes)-1] == "" {
	// 	nodes = nodes[:len(nodes)-1]
	// }
	res := isValidTree(&nodes)
	return res && len(nodes) == 0
}

func isValidTree(nodes *[]string) bool {
	if len(*nodes) == 0 {
		return false
	}
	first := (*nodes)[0]
	*nodes = (*nodes)[1:]

	if first == "#" {
		return true
	}
	left := isValidTree(nodes)
	right := isValidTree(nodes)
	return left && right
}

// @lc code=end

func TestVerifyPreorderSerializationOfABinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "9,3,4,#,#,1,#,#,2,#,6,#,#"\n
// @lcpr case=end

// @lcpr case=start
// "1,#"\n
// @lcpr case=end

// @lcpr case=start
// "9,#,#,1"\n
// @lcpr case=end

*/
