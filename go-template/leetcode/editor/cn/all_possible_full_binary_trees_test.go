/*
 * @lc app=leetcode.cn id=894 lang=golang
 * @lcpr version=30104
 *
 * [894] 所有可能的真二叉树
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
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	return buildFBT(n)

}

func buildFBT(n int) []*TreeNode {
	memo := make(map[int][]*TreeNode)
	var res []*TreeNode

	if n == 1 {
		res = append(res, &TreeNode{Val: 0})
		return res
	}
	if res, found := memo[n]; found {
		return res
	}

	for i := 1; i < n; i += 2 {
		j := n - i - 1
		leftSubTree := buildFBT(i)
		rightSubTree := buildFBT(j)
		for _, left := range leftSubTree {
			for _, right := range rightSubTree {
				root := &TreeNode{Val: 0}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	memo[n] = res
	return res
}

// @lc code=end

func TestAllPossibleFullBinaryTrees(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 7\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
