/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30100
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverseNn(a, k)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseNn(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur, nxt := (*ListNode)(nil), head, head.Next
	for n > 0 {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil {
			nxt = nxt.Next
		}
		n--
	}
	head.Next = cur

	return pre
}

// @lc code=end

func TestReverseNodesInKGroup(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
