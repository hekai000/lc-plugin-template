/*
 * @lc app=leetcode.cn id=662 lang=golang
 * @lcpr version=30104
 *
 * [662] 二叉树最大宽度
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
type Pair struct {
	node *TreeNode
	id   int
}

func widthOfBinaryTree1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 0
	q := []*Pair{{node: root, id: 1}}
	for len(q) > 0 {
		sz := len(q)
		start, end := 0, 0
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			curNode := cur.node
			curId := cur.id
			if i == 0 {
				start = curId
			}
			if i == sz-1 {
				end = curId
			}
			if curNode.Left != nil {
				q = append(q, &Pair{node: curNode.Left, id: curId * 2})
			}
			if curNode.Right != nil {
				q = append(q, &Pair{node: curNode.Right, id: curId*2 + 1})
			}
		}
		maxWidth = mymax11(maxWidth, end-start+1)
	}
	return maxWidth

}
func mymax11(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	firstId := []int{}
	maxWidth := 1
	var traverseWidth func(*TreeNode, int, int)
	traverseWidth = func(root *TreeNode, id int, depth int) {
		if root == nil {
			return
		}

		if depth-1 == len(firstId) {
			firstId = append(firstId, id)
		} else {
			// 计算当前节点到最左侧节点的宽度（含自身），更新最大值
			width := id - firstId[depth-1] + 1
			if width > maxWidth {
				maxWidth = width
			}
		}

		// 递归遍历左右子树（左子树ID=父ID*2，右子树ID=父ID*2+1）
		traverseWidth(root.Left, id*2, depth+1)
		traverseWidth(root.Right, id*2+1, depth+1)
	}
	traverseWidth(root, 1, 1)
	return maxWidth
}

// @lc code=end

func TestMaximumWidthOfBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,3,2,5,3,null,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,5,null,null,9,6,null,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,2,5]\n
// @lcpr case=end

*/
