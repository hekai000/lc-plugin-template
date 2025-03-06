/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30100
 *
 * [23] 合并 K 个升序链表
 */

package leetcode_solutions

import (
	"container/heap"
	"testing"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type PriorityQueueK []*ListNode

func (pq PriorityQueueK) Len() int {
	return len(pq)
}

func (pq PriorityQueueK) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueueK) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueK) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueueK) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{Val: -1}
	p := dummy
	pq := &PriorityQueueK{}
	heap.Init(pq)

	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}

	}

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

// @lc code=end

func TestMergeKSortedLists(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
