/*
 * @lc app=leetcode.cn id=919 lang=golang
 * @lcpr version=30104
 *
 * [919] 完全二叉树插入器
 */

package leetcode_solutions

import (
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
// type CBTInserter struct {
// 	elist []*TreeNode
// 	size  int
// }

// func Constructor(root *TreeNode) CBTInserter {

// 	elist := []*TreeNode{}
// 	size := 0
// 	var bfs func(*TreeNode)
// 	bfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		queue := []*TreeNode{node}

// 		for len(queue) > 0 {
// 			sz := len(queue)
// 			for i := 0; i < sz; i++ {
// 				cur := queue[0]
// 				queue = queue[1:]
// 				elist = append(elist, cur)
// 				size++

// 				if cur.Left != nil {
// 					queue = append(queue, cur.Left)
// 				}
// 				if cur.Right != nil {
// 					queue = append(queue, cur.Right)
// 				}
// 			}
// 		}
// 	}
// 	bfs(root)

// 	return CBTInserter{
// 		elist: elist,
// 		size:  size,
// 	}
// }

// func (this *CBTInserter) Insert(val int) int {
// 	newNode := TreeNode{Val: val}
// 	this.elist = append(this.elist, &newNode)

// 	index := this.size
// 	this.size++
// 	parentIndex := (index - 1) / 2
// 	if index%2 == 1 {
// 		this.elist[parentIndex].Left = &newNode
// 	} else {
// 		this.elist[parentIndex].Right = &newNode
// 	}
// 	return this.elist[parentIndex].Val
// }

// func (this *CBTInserter) Get_root() *TreeNode {
// 	return this.elist[0]
// }

type CBTInserter struct {
	candidate []*TreeNode
	root      *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	candidate := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}

	}

	return CBTInserter{
		candidate: candidate,
		root:      root,
	}
}

func (this *CBTInserter) Insert(val int) int {
	newNode := TreeNode{Val: val}
	node := this.candidate[0]
	if node.Left == nil {
		node.Left = &newNode
	} else {

		node.Right = &newNode
		this.candidate = this.candidate[1:]
	}
	this.candidate = append(this.candidate, &newNode)
	return node.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
// @lc code=end

func TestCompleteBinaryTreeInserter(t *testing.T) {
	// your test code here

}
