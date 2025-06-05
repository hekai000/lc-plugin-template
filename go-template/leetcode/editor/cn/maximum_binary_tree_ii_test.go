/*
 * @lc app=leetcode.cn id=998 lang=golang
 * @lcpr version=30104
 *
 * [998] 最大二叉树 II
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

//分解问题的思维，定义一个函数及其返回值，通过返回值得到原问题的结果
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}

	}
	if val > root.Val {
		temp := root
		root = &TreeNode{Val: val}
		root.Left = temp

	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
	}
	return root
}

// @lc code=end

func TestMaximumBinaryTreeIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,1,3,null,null,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5,2,4,null,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [5,2,3,null,1]\n4\n
// @lcpr case=end

*/
