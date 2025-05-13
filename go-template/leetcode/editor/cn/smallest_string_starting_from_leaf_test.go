/*
 * @lc app=leetcode.cn id=988 lang=golang
 * @lcpr version=30104
 *
 * [988] 从叶结点开始的最小字符串
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
func smallestFromLeaf(root *TreeNode) string {
	var res string
	var path []byte
	trsF(root, &res, &path)
	return res
}

func trsF(root *TreeNode, res *string, path *[]byte) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*path = append(*path, byte('a'+root.Val))
		reverseF(path)
		s := string(*path)
		if *res == "" || *res > s {
			*res = s
		}

		reverseF(path)
		*path = (*path)[:len(*path)-1]
		return
	}
	*path = append(*path, byte('a'+root.Val))
	trsF(root.Left, res, path)
	trsF(root.Right, res, path)
	*path = (*path)[:len(*path)-1]

}

func reverseF(path *[]byte) {
	for i, j := 0, len(*path)-1; i < j; i, j = i+1, j-1 {
		(*path)[i], (*path)[j] = (*path)[j], (*path)[i]
	}
}

// @lc code=end

func TestSmallestStringStartingFromLeaf(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [0,1,2,3,4,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [25,1,3,1,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,null,1,0,null,0]\n
// @lcpr case=end

*/
