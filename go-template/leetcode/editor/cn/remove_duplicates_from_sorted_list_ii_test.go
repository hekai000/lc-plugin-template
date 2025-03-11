/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30100
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates2(head *ListNode) *ListNode {
	// 将原链表分解为两条链表
	// 一条链表存放不重复的节点，另一条链表存放重复的节点
	// 运用虚拟头结点技巧，题目说了 node.val <= 100，所以用 101 作为虚拟头结点
	dummyUniq := &ListNode{Val: 101}
	dummyDup := &ListNode{Val: 101}

	pUniq, pDup := dummyUniq, dummyDup
	p := head

	for p != nil {
		if (p.Next != nil && p.Val == p.Next.Val) || p.Val == pDup.Val {
			// 发现重复节点，接到重复链表后面
			pDup.Next = p
			pDup = pDup.Next
		} else {
			// 不是重复节点，接到不重复链表后面
			pUniq.Next = p
			pUniq = pUniq.Next
		}

		p = p.Next
		// 将原链表和新链表断开
		pUniq.Next = nil
		pDup.Next = nil
	}

	return dummyUniq.Next
}

// @lc code=end

func TestRemoveDuplicatesFromSortedListIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/
