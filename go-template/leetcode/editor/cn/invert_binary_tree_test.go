/*
 * @lc app=leetcode.cn id=226 lang=golang
 * @lcpr version=30104
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	traverseT1(root)
	return root
}

func traverseT1(root *TreeNode) {
	if root == nil {
		return
	}

	traverseT1(root.Left)
	traverseT1(root.Right)
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
}

// @lc code=end

func TestInvertBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,2,7,1,3,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
