/*
 * @lc app=leetcode.cn id=107 lang=golang
 * @lcpr version=30104
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom2(root *TreeNode) [][]int {
	res := levelOrder1(root)
	a := reverseLOB(res)
	return a
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		sz := len(q)
		level := []int{}
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = append([][]int{level}, res...)
	}
	return res
}

func reverseLOB(res [][]int) [][]int {
	len := len(res)
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
func levelOrder1(root *TreeNode) [][]int {
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

func TestBinaryTreeLevelOrderTraversalIi(t *testing.T) {
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
