/*
 * @lc app=leetcode.cn id=1161 lang=golang
 * @lcpr version=30104
 *
 * [1161] 最大层内元素和
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
func maxLevelSum(root *TreeNode) int {
	res := 0

	q := []*TreeNode{root}
	depth := 0
	maxV := math.MinInt32

	for len(q) > 0 {
		levelSum := 0
		depth += 1
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			levelSum += cur.Val

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		if levelSum > maxV {
			res = depth
			maxV = levelSum
		}
	}

	return res
}

// @lc code=end

func TestMaximumLevelSumOfABinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,7,0,7,-8,null,null]\n
// @lcpr case=end

// @lcpr case=start
// [989,null,10250,98693,-89388,null,null,null,-32127]\n
// @lcpr case=end

*/
