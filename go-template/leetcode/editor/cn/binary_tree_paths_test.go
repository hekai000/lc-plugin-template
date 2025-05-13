/*
 * @lc app=leetcode.cn id=257 lang=golang
 * @lcpr version=30104
 *
 * [257] 二叉树的所有路径
 */

package leetcode_solutions

import (
	"strconv"
	"strings"
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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	path := []string{}
	travereBTP(root, &res, &path)
	return res
}

func travereBTP(root *TreeNode, res *[]string, path *[]string) {
	if root == nil {
		return
	}
	*path = append(*path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(*path, "->"))
		*path = (*path)[:len(*path)-1]
		return

	}
	travereBTP(root.Left, res, path)
	travereBTP(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
	return

}

// @lc code=end

func TestBinaryTreePaths(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
