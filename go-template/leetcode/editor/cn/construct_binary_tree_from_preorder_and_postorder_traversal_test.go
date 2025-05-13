/*
 * @lc app=leetcode.cn id=889 lang=golang
 * @lcpr version=30104
 *
 * [889] 根据前序和后序遍历构造二叉树
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
var valToIndexPost map[int]int

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	valToIndexPost = make(map[int]int)
	for index, val := range postorder {
		valToIndexPost[val] = index
	}

	return buildPP(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func buildPP(preorder []int, preStart int, preEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	if preStart == preEnd {
		return &TreeNode{
			Val:   preorder[preStart],
			Left:  nil,
			Right: nil,
		}
	}
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	leftRootVal := preorder[preStart+1]
	leftIndex := valToIndexPost[leftRootVal]
	leftSize := leftIndex - postStart + 1
	root.Left = buildPP(preorder, preStart+1, preStart+leftSize, postorder, postStart, leftIndex)
	root.Right = buildPP(preorder, preStart+leftSize+1, preEnd, postorder, leftIndex+1, postEnd-1)
	return root
}

// @lc code=end

func TestConstructBinaryTreeFromPreorderAndPostorderTraversal(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n[1]\n
// @lcpr case=end

*/
