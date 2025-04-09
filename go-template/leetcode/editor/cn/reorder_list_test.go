/*
 * @lc app=leetcode.cn id=143 lang=golang
 * @lcpr version=30104
 *
 * [143] 重排链表
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
func reorderList(head *ListNode) {
	stk := []*ListNode{}
	p := head
	for p != nil {
		stk = append(stk, p)
		p = p.Next
	}
	p = head
	for p != nil {
		lastNode := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		nextNode := p.Next
		if lastNode == nextNode || lastNode.Next == nextNode {
			lastNode.Next = nil
			break
		}
		p.Next = lastNode
		lastNode.Next = nextNode
		p = nextNode
	}
}

// @lc code=end

func TestReorderList(t *testing.T) {
	// your test code here
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	reorderList(&node1)

}

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/
