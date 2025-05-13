/*
 * @lc app=leetcode.cn id=1022 lang=golang
 * @lcpr version=30104
 *
 * [1022] 从根到叶的二进制数之和
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
func sumRootToLeaf(root *TreeNode) int {
	var res int
	var path []int

	trSR(root, &res, &path)
	return res
}

func trSR(root *TreeNode, res *int, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		val := transFpath(path)
		*res += val
		*path = (*path)[:len(*path)-1]
		return
	}
	trSR(root.Left, res, path)
	trSR(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
}

func transFpath(path *[]int) int {
	sum := 0
	for i := 0; i < len(*path); i++ {
		sum = sum*2 + (*path)[i]
	}
	return sum
}

// @lc code=end

func TestSumOfRootToLeafBinaryNumbers(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,0,1,0,1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
