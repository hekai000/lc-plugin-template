/*
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=30104
 *
 * [104] 二叉树的最大深度
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
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return Mymax1(leftMax, rightMax) + 1
}

func maxDepth(root *TreeNode) int {

	depth := 0
	res := 0
	backtrack(root, &depth, &res)
	return res
}

func backtrack(root *TreeNode, depth *int, res *int) {
	if root == nil {
		return
	}
	*depth++
	if root.Left == nil && root.Right == nil {
		*res = Mymax1(*res, *depth)
	}
	backtrack(root.Left, depth, res)
	backtrack(root.Right, depth, res)
	*depth--
}
func Mymax1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2]\n
// @lcpr case=end

*/
