/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30100
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

func TestAddTwoNumbers(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/
