/*
 * @lc app=leetcode.cn id=958 lang=golang
 * @lcpr version=30104
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	end := false
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur == nil {
				end = true
			} else {
				if end {
					return false
				}
				q = append(q, cur.Left)
				q = append(q, cur.Right)
			}
		}
	}
	return true
}

// @lc code=end

func TestCheckCompletenessOfABinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

*/
