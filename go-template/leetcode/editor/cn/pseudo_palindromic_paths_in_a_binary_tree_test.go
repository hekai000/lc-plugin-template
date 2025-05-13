/*
 * @lc app=leetcode.cn id=1457 lang=golang
 * @lcpr version=30104
 *
 * [1457] 二叉树中的伪回文路径
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
// func pseudoPalindromicPaths(root *TreeNode) int {
// 	var res int
// 	count := make([]int, 10)

// 	trPPP2(root, count, &res)

// 	return res

// }
func pseudoPalindromicPaths(root *TreeNode) int {
	var res int
	var count int

	trPPP3(root, &count, &res)

	return res

}

func trPPP3(root *TreeNode, count *int, res *int) {
	if root == nil {
		return
	}
	*count ^= 1 << root.Val
	if root.Left == nil && root.Right == nil {
		if *count&(*count-1) == 0 {
			*res++

		}
		*count ^= 1 << root.Val
		return
	}
	trPPP3(root.Left, count, res)
	trPPP3(root.Right, count, res)
	*count ^= 1 << root.Val
	return
}

func trPPP2(root *TreeNode, count []int, res *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		// 遇到叶子节点，判断路径是否是伪回文串
		count[root.Val]++
		// 如果路径上出现奇数次的数字个数大于 1，
		// 则不可能组成回文串，反之则可以组成回文串
		odd := 0
		for _, n := range count {
			if n%2 == 1 {
				odd++
			}
		}
		//出现奇数次的元素个数小于等于1,
		if odd <= 1 {
			*res++
		}
		count[root.Val]--
		return
	}

	count[root.Val]++
	// 二叉树遍历框架
	trPPP2(root.Left, count, res)
	trPPP2(root.Right, count, res)

	count[root.Val]--
}

// func pseudoPalindromicPaths(root *TreeNode) int {
// 	var res [][]int
// 	var path []int
// 	trPPP(root, &res, &path)
// 	c := countPalind(res)
// 	return c

// }

func trPPP(root *TreeNode, res *[][]int, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		newPath := make([]int, len(*path))
		copy(newPath, *path)
		*res = append(*res, newPath)
		*path = (*path)[:len(*path)-1]
		return
	}

	trPPP(root.Left, res, path)
	trPPP(root.Right, res, path)
	*path = (*path)[:len(*path)-1]
	return
}

func countPalind(res [][]int) int {
	count := 0
	for _, v := range res {
		if canPermutePalindrome(v) {
			count++
		}
	}
	return count
}

func canPermutePalindrome(arr []int) bool {
	// 统计元素频率
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	// 统计奇数次元素的个数
	oddCount := 0
	for _, count := range freq {
		if count%2 != 0 {
			oddCount++
		}
	}

	// 根据数组长度判断条件
	if len(arr)%2 == 0 {
		return oddCount == 0
	} else {
		return oddCount <= 1
	}
}

// @lc code=end

func TestPseudoPalindromicPathsInABinaryTree(t *testing.T) {
	// your test code here
	arr := []interface{}{2, 3, 1, 3, 1, nil, 1}
	root := CreateRoot(arr)
	pseudoPalindromicPaths(root)

}

/*
// @lcpr case=start
// [2,3,1,3,1,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,1,1,3,null,null,null,null,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [9]\n
// @lcpr case=end

*/
