/*
 * @lc app=leetcode.cn id=230 lang=golang
 * @lcpr version=30104
 *
 * [230] 二叉搜索树中第 K 小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	res := []int{}
	traverseKS(root, &res)
	return res[k-1]
}

// 遍历二叉树，将节点值存储在res中
func traverseKS(root *TreeNode, res *[]int) {
	// 如果当前节点为空，则返回
	if root == nil {
		return
	}
	// 递归遍历左子树
	traverseKS((*root).Left, res)
	*res = append(*res, root.Val)
	traverseKS((*root).Right, res)
}

// @lc code=end

func TestKthSmallestElementInABst(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,1,4,null,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,null,1]\n3\n
// @lcpr case=end

*/
