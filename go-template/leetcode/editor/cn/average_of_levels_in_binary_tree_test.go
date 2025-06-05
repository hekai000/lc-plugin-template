/*
 * @lc app=leetcode.cn id=637 lang=golang
 * @lcpr version=30104
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {

	res := []float64{}
	if root == nil {
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	depth := 0
	for len(q) > 0 {
		sz := len(q)
		depth += 1
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if len(res) == depth-1 {
				res = append(res, float64(cur.Val))
			} else {
				res[depth-1] = res[depth-1] + float64(cur.Val)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res[depth-1] = res[depth-1] / float64(sz)
	}
	return res
}

// @lc code=end

func TestAverageOfLevelsInBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [3,9,20,15,7]\n
// @lcpr case=end

*/
