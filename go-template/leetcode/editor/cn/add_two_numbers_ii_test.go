/*
 * @lc app=leetcode.cn id=445 lang=golang
 * @lcpr version=30100
 *
 * [445] 两数相加 II
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

func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	var stk1 []int
	var stk2 []int
	for l1 != nil {
		stk1 = append(stk1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stk2 = append(stk2, l2.Val)
		l2 = l2.Next
	}

	carry := 0

	dummyNode := &ListNode{Val: -1}

	for len(stk1) > 0 || len(stk2) > 0 || carry > 0 {
		value := carry
		if len(stk1) > 0 {
			value += stk1[len(stk1)-1]
			stk1 = stk1[:(len(stk1) - 1)]
		}
		if len(stk2) > 0 {
			value += stk2[len(stk2)-1]
			stk2 = stk2[:(len(stk2) - 1)]
		}
		carry = value / 10
		value = value % 10
		newNode := &ListNode{Val: value, Next: dummyNode.Next}
		dummyNode.Next = newNode

	}
	return dummyNode.Next
}
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	l11 := reverseList(l1)
	l22 := reverseList(l2)
	l3 := addTwoNumbers2(l11, l22)
	return reverseList(l3)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, cur, next *ListNode
	prev, cur, next = nil, head, head.Next
	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	return prev
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: -1}
	p1 := l1
	p2 := l2
	cur := dummyNode
	carry := 0
	for p1 != nil || p2 != nil || carry > 0 {
		value := carry
		if p1 != nil {
			value += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			value += p2.Val
			p2 = p2.Next
		}
		carry = value / 10
		value = value % 10

		newNode := &ListNode{Val: value}
		cur.Next = newNode
		cur = cur.Next

	}

	return dummyNode.Next
}

// @lc code=end

func TestAddTwoNumbersIi(t *testing.T) {
	// your test code here
	l1 := CreateHead([]int{1, 2, 4})
	l2 := CreateHead([]int{1, 3, 4})
	result := addTwoNumbers4(l1, l2)
	PrintList(result)

}

/*
// @lcpr case=start
// [7,2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

*/
