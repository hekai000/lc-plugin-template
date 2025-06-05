/*
 * @lc app=leetcode.cn id=1110 lang=golang
 * @lcpr version=30104
 *
 * [1110] 删点成林
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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delSet := make(map[int]bool)
	for _, v := range to_delete {
		delSet[v] = true
	}
	res := make([]*TreeNode, 0)
	doDelNodes(root, false, &res, delSet)
	return res
}
func doDelNodes(root *TreeNode, hasParent bool, res *[]*TreeNode, delSet map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}
	deleted := false
	if _, ok := delSet[root.Val]; ok {
		deleted = true
	}
	if !deleted && !hasParent {
		*res = append(*res, root)
	}
	root.Left = doDelNodes(root.Left, !deleted, res, delSet)
	root.Right = doDelNodes(root.Right, !deleted, res, delSet)

	if deleted {
		return nil
	}

	return root
}

func doDelete(root *TreeNode, hasParent bool, res *[]*TreeNode, delSet map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}

	deleted := delSet[root.Val]
	if !deleted && !hasParent {
		*res = append(*res, root)
	}
	root.Left = doDelete(root.Left, !deleted, res, delSet)
	root.Right = doDelete(root.Right, !deleted, res, delSet)

	if deleted {
		return nil
	}
	return root
}

// @lc code=end

func TestDeleteNodesAndReturnForest(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n[3,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,4,null,3]\n[3]\n
// @lcpr case=end

*/
