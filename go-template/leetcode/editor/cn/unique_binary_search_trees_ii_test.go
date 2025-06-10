/*
 * @lc app=leetcode.cn id=95 lang=golang
 * @lcpr version=30104
 *
 * [95] 不同的二叉搜索树 II
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
//定义buildBST为返回start，end区间内所有有效的BST。
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return buildBST(1, n)
}

func buildBST(start, end int) []*TreeNode {
	res := []*TreeNode{}
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftTrees := buildBST(start, i-1)
		rightTrees := buildBST(i+1, end)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = leftTree
				root.Right = rightTree
				res = append(res, root)
			}
		}
	}
	return res
}

// @lc code=end

func TestUniqueBinarySearchTreesIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
