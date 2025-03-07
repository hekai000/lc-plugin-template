/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30100
 *
 * [92] 反转链表 II
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

//递归写法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var successor *ListNode

	var reverseNr func(head *ListNode, n int) *ListNode
	reverseNr = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			successor = head.Next
			return head
		}
		last := reverseNr(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return last
	}

	if left == 1 {
		return reverseNr(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head

}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN2(head, right)
	}
	prev := head
	for i := 1; i < left-1; i++ {
		prev = prev.Next
	}
	prev.Next = reverseN2(prev.Next, right-left+1)
	return head
}

func reverseN2(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *ListNode
	prev, cur, next = nil, head, head.Next

	for n > 0 {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
		n--
	}
	head.Next = cur
	return prev
}

// @lc code=end

func TestReverseLinkedListIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
