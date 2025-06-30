/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=30104
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{}
	q = append(q, root)
	for len(q) > 0 {
		sz := len(q)
		var pre *Node
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return root
}

// @lc code=end

func TestPopulatingNextRightPointersInEachNodeIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
