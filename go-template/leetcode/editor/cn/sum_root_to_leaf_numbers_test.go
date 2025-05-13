/*
 * @lc app=leetcode.cn id=129 lang=golang
 * @lcpr version=30104
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	path := []int{}
	res := []int{}
	traSumN(root, &path, &res)

	sumi := 0
	for i := 0; i < len(res); i++ {
		sumi += res[i]
	}
	return sumi
}

func traSumN(root *TreeNode, path *[]int, res *[]int) {
	if root == nil {
		return
	}

	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		sum := 0
		for i := 0; i < len(*path); i++ {
			sum = sum*10 + (*path)[i]

		}
		*res = append(*res, sum)
		*path = (*path)[:len(*path)-1]
		return
	}
	traSumN(root.Left, path, res)
	traSumN(root.Right, path, res)
	*path = (*path)[:len(*path)-1]
	return
}

// @lc code=end

func TestSumRootToLeafNumbers(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,0,5,1]\n
// @lcpr case=end

*/
