/*
 * @lc app=leetcode.cn id=103 lang=golang
 * @lcpr version=30104
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	flag := 1
	for len(q) > 0 {
		sz := len(q)
		level := []int{}
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if flag == 1 {
				level = append(level, cur.Val)
			} else {
				level = append([]int{cur.Val}, level...)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		flag = -flag
		res = append(res, level)
	}
	return res

}

// @lc code=end

func TestBinaryTreeZigzagLevelOrderTraversal(t *testing.T) {
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
