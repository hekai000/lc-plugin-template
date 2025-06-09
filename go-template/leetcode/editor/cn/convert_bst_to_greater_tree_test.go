/*
 * @lc app=leetcode.cn id=538 lang=golang
 * @lcpr version=30104
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var traverseCBST func(*TreeNode)
	traverseCBST = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverseCBST(root.Right)
		sum += root.Val
		root.Val = sum
		traverseCBST(root.Left)
	}
	traverseCBST(root)
	return root

}

// @lc code=end

func TestConvertBstToGreaterTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]\n
// @lcpr case=end

// @lcpr case=start
// [0,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4,1]\n
// @lcpr case=end

*/
