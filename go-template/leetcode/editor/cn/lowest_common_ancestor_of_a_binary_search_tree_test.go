/*
 * @lc app=leetcode.cn id=235 lang=golang
 * @lcpr version=30104
 *
 * [235] 二叉搜索树的最近公共祖先
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca := findlca(root, p, q)
	return lca
}

func findlca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := findlca(root.Left, p, q)
	right := findlca(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// @lc code=end

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n8\n
// @lcpr case=end

// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n4\n
// @lcpr case=end

*/
