/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=30104
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	p := root

	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
	return
}

// @lc code=end

func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
