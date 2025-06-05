/*
 * @lc app=leetcode.cn id=102 lang=golang
 * @lcpr version=30104
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{}
	res := [][]int{}
	q = append(q, root)

	for len(q) > 0 {
		sz := len(q)
		level := make([]int, sz)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			level[i] = cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// @lc code=end

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
