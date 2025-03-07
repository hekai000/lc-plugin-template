/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30100
 *
 * [234] 回文链表
 */

package leetcode_solutions

import (
	"fmt"
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
func isPalindrome2(head *ListNode) bool {
	left := head
	var right *ListNode
	var traverse func(right *ListNode)
	res := true
	traverse = func(right *ListNode) {
		if right == nil {
			return
		}

		traverse(right.Next)
		if right.Val != left.Val {
			res = false
		}
		left = left.Next
	}

	right = head
	traverse(right)
	return res

}
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// if fast != nil {
	// 	slow = slow.Next
	// }

	right := reverseList2(slow)
	left := head
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}

	return true
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	prev, cur := nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next

	}
	return prev
}

// @lc code=end

func TestPalindromeLinkedList(t *testing.T) {
	// your test code here
	p := CreateHead([]int{1, 2, 2, 1})

	a := isPalindrome(p)
	fmt.Println(a)
}

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
