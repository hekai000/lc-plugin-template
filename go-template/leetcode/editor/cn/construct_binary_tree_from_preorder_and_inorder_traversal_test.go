/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=30104
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
// var valToIndex map[int]int

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	valToIndex = make(map[int]int)
// 	for index, val := range inorder {
// 		valToIndex[val] = index
// 	}
// 	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)

// }

// func build(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
// 	if preStart > preEnd {
// 		return nil
// 	}
// 	rootVal := preorder[preStart]
// 	root := &TreeNode{Val: rootVal}

// 	index := valToIndex[rootVal]

// 	leftSize := index - inStart
// 	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, index-1)
// 	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, index+1, inEnd)
// 	return root
// }

// // @lc code=end

// func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {
// 	// your test code here

// }

// /*
// // @lcpr case=start
// // [3,9,20,15,7]\n[9,3,15,20,7]\n
// // @lcpr case=end

// // @lcpr case=start
// // [-1]\n[-1]\n
// // @lcpr case=end

// */
