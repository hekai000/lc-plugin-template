/*
 * @lc app=leetcode.cn id=515 lang=golang
 * @lcpr version=30104
 *
 * [515] 在每个树行中找最大值
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
// func largestValues(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}
// 	q := []*TreeNode{}
// 	q = append(q, root)
// 	dept := 0
// 	for len(q) > 0 {
// 		sz := len(q)
// 		dept += 1
// 		for i := 0; i < sz; i++ {
// 			cur := q[0]
// 			q = q[1:]
// 			if len(res) == dept-1 {
// 				res = append(res, cur.Val)
// 			} else {
// 				if cur.Val > res[dept-1] {
// 					res[dept-1] = cur.Val
// 				}
// 			}
// 			if cur.Left != nil {
// 				q = append(q, cur.Left)
// 			}
// 			if cur.Right != nil {
// 				q = append(q, cur.Right)
// 			}
// 		}
// 	}
// 	return res
// }
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var traverseLV func(*TreeNode, int)
	traverseLV = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth-1 {
			res = append(res, root.Val)
		} else {
			if root.Val > res[depth-1] {
				res[depth-1] = root.Val
			}
		}
		traverseLV(root.Left, depth+1)
		traverseLV(root.Right, depth+1)
	}
	traverseLV(root, 1)
	return res
}

// @lc code=end

func TestFindLargestValueInEachTreeRow(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,3,2,5,3,null,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
