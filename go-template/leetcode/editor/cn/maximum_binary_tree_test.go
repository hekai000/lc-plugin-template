/*
 * @lc app=leetcode.cn id=654 lang=golang
 * @lcpr version=30104
 *
 * [654] 最大二叉树
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

func getMaxIndex(nums []int) int {
	max := math.MinInt32
	index := -1
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index

}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	//buildMaxTree(nums, index),以index为根节点，构建最大二叉树，返回根节点

	root := buildMaxTree(nums)
	return root
}

func buildMaxTree(nums []int) *TreeNode {
	maxIndex := getMaxIndex(nums)
	if maxIndex == -1 {
		return nil
	}
	root := &TreeNode{Val: nums[maxIndex]}
	leftIndex := 0

	root.Left = buildMaxTree(nums[leftIndex:maxIndex])
	root.Right = buildMaxTree(nums[maxIndex+1:])
	return root

}

// @lc code=end

func TestMaximumBinaryTree(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,2,1,6,0,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

*/
