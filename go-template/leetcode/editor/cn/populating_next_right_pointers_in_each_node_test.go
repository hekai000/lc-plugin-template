/*
 * @lc app=leetcode.cn id=116 lang=golang
 * @lcpr version=30104
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

package leetcode_solutions

// import "testing"

// // @lc code=start
// /**
//  * Definition for a Node.
//  * type Node struct {
//  *     Val int
//  *     Left *Node
//  *     Right *Node
//  *     Next *Node
//  * }
//  */

// func connect1(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	traversec(root.Left, root.Right)
// 	return root
// }

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	traversec1(root)
// 	return root
// }

// func traversec1(node *Node) {
// 	if node == nil || node.Left == nil {
// 		return
// 	}

// 	node.Left.Next = node.Right
// 	if node.Next != nil {
// 		node.Right.Next = node.Next.Left
// 	}
// 	traversec1(node.Left)
// 	traversec1(node.Right)
// 	return

// }

// func traversec(node1 *Node, node2 *Node) {
// 	if node1 == nil || node2 == nil {
// 		return
// 	}

// 	node1.Next = node2
// 	traversec(node1.Left, node1.Right)
// 	traversec(node2.Left, node2.Right)
// 	traversec(node1.Right, node2.Left)

// }

// // @lc code=end

// func TestPopulatingNextRightPointersInEachNode(t *testing.T) {
// 	// your test code here

// }

// /*
// // @lcpr case=start
// // [1,2,3,4,5,6,7]\n
// // @lcpr case=end

// // @lcpr case=start
// // []\n
// // @lcpr case=end

// */
