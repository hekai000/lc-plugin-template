/*
 * @lc app=leetcode.cn id=297 lang=golang
 * @lcpr version=30104
 *
 * [297] 二叉树的序列化与反序列化
 */

package leetcode_solutions

import (
	"container/list"
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

// type Codec struct {
// 	SEP  string
// 	NULL string
// }

// func Constructor() Codec {
// 	return Codec{
// 		SEP:  ",",
// 		NULL: "#",
// 	}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	var sb strings.Builder
// 	this._sericalize(root, &sb)
// 	return sb.String()
// }

// func (this *Codec) _sericalize(root *TreeNode, sb *strings.Builder) {
// 	if root == nil {
// 		sb.WriteString(this.NULL)
// 		sb.WriteString(this.SEP)
// 		return
// 	}
// 	sb.WriteString(strconv.Itoa(root.Val))
// 	sb.WriteString(this.SEP)

// 	this._sericalize(root.Left, sb)
// 	this._sericalize(root.Right, sb)
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	// nodes := make([]string, 0)
// 	// for _, s := range strings.Split(data, this.SEP) {
// 	// 	nodes = append(nodes, s)
// 	// }
// 	nodes := strings.Split(data, this.SEP)
// 	if nodes[len(nodes)-1] == "" {
// 		nodes = nodes[:len(nodes)-1]
// 	}
// 	return this._deserialize(&nodes)
// }

// func (this *Codec) _deserialize(nodes *[]string) *TreeNode {
// 	if len(*nodes) == 0 {
// 		return nil
// 	}

// 	first := (*nodes)[0]
// 	*nodes = (*nodes)[1:]

// 	if first == this.NULL {
// 		return nil
// 	}

// 	val, err := strconv.Atoi(first)
// 	if err != nil {
// 		panic("Parse Error")
// 	}
// 	root := &TreeNode{Val: val}
// 	root.Left = this._deserialize(nodes)
// 	root.Right = this._deserialize(nodes)
// 	return root

// }

// type Codec struct {
// 	SEP  string
// 	NULL string
// }

// func Constructor() Codec {
// 	return Codec{
// 		SEP:  ",",
// 		NULL: "#",
// 	}
// }

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	var sb strings.Builder
// 	this._sericalize(root, &sb)
// 	return sb.String()
// }

// func (this *Codec) _sericalize(root *TreeNode, sb *strings.Builder) {
// 	if root == nil {
// 		sb.WriteString(this.NULL)
// 		sb.WriteString(this.SEP)
// 		return
// 	}

// 	this._sericalize(root.Left, sb)
// 	this._sericalize(root.Right, sb)
// 	sb.WriteString(strconv.Itoa(root.Val))
// 	sb.WriteString(this.SEP)

// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	nodes := strings.Split(data, this.SEP)
// 	if nodes[len(nodes)-1] == "" {
// 		nodes = nodes[:len(nodes)-1]
// 	}
// 	return this._deserialize(&nodes)
// }

// func (this *Codec) _deserialize(nodes *[]string) *TreeNode {
// 	if len(*nodes) == 0 {
// 		return nil
// 	}

// 	last := (*nodes)[len(*nodes)-1]
// 	*nodes = (*nodes)[:len(*nodes)-1]

// 	if last == this.NULL {
// 		return nil
// 	}

// 	val, err := strconv.Atoi(last)
// 	if err != nil {
// 		panic("Parse Error")
// 	}
// 	root := &TreeNode{Val: val}
// 	root.Right = this._deserialize(nodes)
// 	root.Left = this._deserialize(nodes)

// 	return root

// }

type Codec struct {
	SEP  string
	NULL string
}

func Constructor00() Codec {
	return Codec{
		SEP:  ",",
		NULL: "#",
	}
}

// Serializes a tree to a single string.

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {

		return ""
	}
	var sb strings.Builder
	q := list.New()
	q.PushBack(root)
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			cur := q.Remove(q.Front()).(*TreeNode)
			if cur == nil {
				sb.WriteString(this.NULL)
				sb.WriteString(this.SEP)
				continue
			}
			sb.WriteString(strconv.Itoa(cur.Val))
			sb.WriteString(this.SEP)

			q.PushBack(cur.Left)
			q.PushBack(cur.Right)

		}
	}
	return sb.String()

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize2(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data, this.SEP)
	rootVal, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: rootVal}

	queue := list.New()
	queue.PushBack(root)
	index := 1
	for queue.Len() > 0 {
		parent := queue.Remove(queue.Front()).(*TreeNode)
		leftValStr := nodes[index]
		index++

		if leftValStr != this.NULL {
			leftVal, _ := strconv.Atoi(leftValStr)
			parent.Left = &TreeNode{Val: leftVal}
			queue.PushBack(parent.Left)
		}

		rightValstr := nodes[index]
		index++
		if rightValstr != this.NULL {
			rightVal, _ := strconv.Atoi(rightValstr)
			parent.Right = &TreeNode{Val: rightVal}
			queue.PushBack(parent.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

func TestSerializeAndDeserializeBinaryTree(t *testing.T) {
	// your test code here
	data := "1,2,3,#,4,#,#,#,#"
	ser := Constructor00()
	ser.deserialize2(data)
}

/*
// @lcpr case=start
// [1,2,3,null,null,4,5]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
