/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=30104
 *
 * [199] 二叉树的右视图
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	//BFS
	// if root == nil {
	// 	return []int{}
	// }
	// var res []int
	// q := []*TreeNode{root}

	// for len(q) > 0 {
	// 	sz := len(q)
	// 	last := q[0].Val
	// 	for i := 0; i < sz; i++ {
	// 		cur := q[0]
	// 		q = q[1:]
	// 		if cur.Right != nil {
	// 			q = append(q, cur.Right)
	// 		}
	// 		if cur.Left != nil {
	// 			q = append(q, cur.Left)
	// 		}
	// 	}
	// 	res = append(res, last)
	// }
	// return res
	var res []int
	trRightSideView(root, 0, &res)
	return res
}

func trRightSideView(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}

	if len(*res) == depth {
		*res = append(*res, root.Val)
	}
	trRightSideView(root.Right, depth+1, res)
	trRightSideView(root.Left, depth+1, res)
	return
}

// @lc code=end

func TestBinaryTreeRightSideView(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,null,null,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
