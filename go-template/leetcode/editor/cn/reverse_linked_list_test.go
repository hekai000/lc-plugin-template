/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=30100
 *
 * [206] 反转链表
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList3(head *ListNode) *ListNode {
	var prev, cur *ListNode

	prev, cur = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

//递归写法，主要是要明确递归函数的意义
func reverseList4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList4(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// @lc code=end

func TestReverseLinkedList(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
