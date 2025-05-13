/*
 * @lc app=leetcode.cn id=144 lang=golang
 * @lcpr version=30104
 *
 * [144] 二叉树的前序遍历
 */

package leetcode_solutions

// import "testing"

// // @lc code=start
// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// func preorderTraversal2(root *TreeNode) []int {

// 	res := []int{}

// 	var traverse func(root *TreeNode)
// 	traverse = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		res = append(res, root.Val)
// 		traverse(root.Left)
// 		traverse(root.Right)
// 	}
// 	traverse(root)
// 	return res
// }

// func preorderTraversal(root *TreeNode) []int {

// 	var res []int
// 	if root == nil {
// 		return res
// 	}

// 	res = append(res, root.Val)
// 	res = append(res, preorderTraversal(root.Left)...)
// 	res = append(res, preorderTraversal(root.Right)...)
// 	return res
// }

// // @lc code=end

// func TestBinaryTreePreorderTraversal(t *testing.T) {
// 	// your test code here

// }

// /*
// // @lcpr case=start
// // [1,null,2,3]\n
// // @lcpr case=end

// // @lcpr case=start
// // [1,2,3,4,5,null,8,null,null,6,7,9]\n
// // @lcpr case=end

// // @lcpr case=start
// // []\n
// // @lcpr case=end

// // @lcpr case=start
// // [1]\n
// // @lcpr case=end

// */
