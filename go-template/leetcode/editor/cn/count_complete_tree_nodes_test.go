/*
 * @lc app=leetcode.cn id=222 lang=golang
 * @lcpr version=30104
 *
 * [222] 完全二叉树的节点个数
 */

package leetcode_solutions

import (
	"math"
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
func countNodes(root *TreeNode) int {
	hl, hr := 0, 0

	l, r := root, root
	for l != nil {
		l = l.Left
		hl++
	}

	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr {
		return int(math.Pow(2, float64(hl))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// @lc code=end

func TestCountCompleteTreeNodes(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
