/*
 * @lc app=leetcode.cn id=701 lang=golang
 * @lcpr version=30104
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

// @lc code=end

func TestInsertIntoABinarySearchTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [40,20,60,10,30,50,70]\n25\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3,null,null,null,null,null,null]\n5\n
// @lcpr case=end

*/
