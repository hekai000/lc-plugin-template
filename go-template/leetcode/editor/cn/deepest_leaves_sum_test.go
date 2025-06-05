/*
 * @lc app=leetcode.cn id=1302 lang=golang
 * @lcpr version=30104
 *
 * [1302] 层数最深叶子节点的和
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
func deepestLeavesSum(root *TreeNode) int {

	res := []int{}

	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		sz := len(q)
		depth += 1
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if len(res) == depth-1 {
				res = append(res, cur.Val)
			} else {
				res[depth-1] += cur.Val
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return res[depth-1]

}

// @lc code=end

func TestDeepestLeavesSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,null,6,7,null,null,null,null,8]\n
// @lcpr case=end

// @lcpr case=start
// [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]\n
// @lcpr case=end

*/
