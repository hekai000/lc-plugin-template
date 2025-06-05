/*
 * @lc app=leetcode.cn id=1609 lang=golang
 * @lcpr version=30104
 *
 * [1609] 奇偶树
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
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root}

	even := true
	for len(q) > 0 {
		sz := len(q)
		var prev int
		if even {
			prev = math.MinInt32
		} else {
			prev = math.MaxInt32
		}
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			if even {
				if prev >= cur.Val || cur.Val%2 == 0 {
					return false
				}
			} else {
				if prev <= cur.Val || cur.Val%2 == 1 {
					return false
				}
			}
			prev = cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}

		}
		even = !even
	}
	return true
}

// @lc code=end

func TestEvenOddTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,10,4,3,null,7,9,12,8,6,null,null,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,2,3,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [5,9,1,3,5,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [11,8,6,1,3,9,11,30,20,18,16,12,10,4,2,17]\n
// @lcpr case=end

*/
