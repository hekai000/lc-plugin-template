/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=30104
 *
 * [98] 验证二叉搜索树
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isValidBST(root *TreeNode) bool {
// 	min, max := math.MinInt64, math.MaxInt64
// 	if root == nil {
// 		return true
// 	}
// 	return isValidBSTHelper(root, min, max)

// }
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !dfs(node.Left) {
			return false
		}

		if node.Val <= pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)

}
func isValidBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}

// @lc code=end

func TestValidateBinarySearchTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
