/*
 * @lc app=leetcode.cn id=382 lang=golang
 * @lcpr version=30104
 *
 * [382] 链表随机节点
 */

package leetcode_solutions

import (
	"math/rand"
	"testing"
	"time"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head      *ListNode
	generator *rand.Rand
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head:      head,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) GetRandom() int {
	i := 0
	p := this.head
	res := 0
	for p != nil {
		i++
		if 0 == this.generator.Intn(i) {
			res = p.Val
		}
		p = p.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

func TestLinkedListRandomNode(t *testing.T) {
	// your test code here

}
