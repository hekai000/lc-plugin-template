/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=30104
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

// var valToIndexP map[int]int

// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	valToIndexP = make(map[int]int)
// 	for index, value := range inorder {
// 		valToIndexP[value] = index
// 	}

// 	return buildF(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)

// }

// func buildF(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
// 	if postEnd < postStart {
// 		return nil
// 	}
// 	rootVal := postorder[postEnd]
// 	index := valToIndexP[rootVal]
// 	root := &TreeNode{Val: rootVal}

// 	leftSize := index - inStart

// 	root.Left = buildF(inorder, inStart, index-1, postorder, postStart, postStart+leftSize-1)
// 	root.Right = buildF(inorder, index+1, inEnd, postorder, postStart+leftSize, postEnd-1)
// 	return root

// }

// // @lc code=end

// func TestConstructBinaryTreeFromInorderAndPostorderTraversal(t *testing.T) {
// 	// your test code here

// }

// /*
// // @lcpr case=start
// // [9,3,15,20,7]\n[9,15,7,20,3]\n
// // @lcpr case=end

// // @lcpr case=start
// // [-1]\n[-1]\n
// // @lcpr case=end

// */
