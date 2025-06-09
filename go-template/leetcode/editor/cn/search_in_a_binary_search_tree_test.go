/*
 * @lc app=leetcode.cn id=700 lang=golang
 * @lcpr version=30104
 *
 * [700] 二叉搜索树中的搜索
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return root
}

// @lc code=end

func TestSearchInABinarySearchTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,2,7,1,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

*/
