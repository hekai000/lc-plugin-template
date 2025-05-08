/*
 * @lc app=leetcode.cn id=543 lang=golang
 * @lcpr version=30104
 *
 * [543] 二叉树的直径
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

func diameterOfBinaryTree(root *TreeNode) int {
	maxDia := 0

	var maxDep func(root *TreeNode) int
	maxDep = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftM := maxDep(root.Left)
		rightM := maxDep(root.Right)
		maxDia = max2(leftM+rightM, maxDia)
		return 1 + max2(leftM, rightM)

	}
	maxDep(root)
	return maxDia
}
func diameterOfBinaryTre2(root *TreeNode) int {
	maxDiameter := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		leftMax := maxDepth2(root.Left)
		rightMax := maxDepth2(root.Right)
		myDiameter := leftMax + rightMax
		maxDiameter = max2(myDiameter, maxDiameter)

		traverse(root.Left)
		traverse(root.Right)

	}

	traverse(root)
	return maxDiameter

}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	return max2(left+1, right+1)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestDiameterOfBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
