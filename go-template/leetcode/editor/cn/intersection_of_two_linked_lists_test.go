/*
 * @lc app=leetcode.cn id=160 lang=golang
 * @lcpr version=30100
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// p1 指向 A 链表头结点，p2 指向 B 链表头结点
	p1 := headA
	p2 := headB
	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// @lc code=end

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/
