/*
 * @lc app=leetcode.cn id=1019 lang=golang
 * @lcpr version=30104
 *
 * [1019] 链表中的下一个更大节点
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
func nextLargerNodes(head *ListNode) []int {
	arr := []int{}
	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}
	n := len(arr)
	res := make([]int, n)
	stk := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] <= arr[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = 0
		} else {
			res[i] = stk[len(stk)-1]
		}
		stk = append(stk, arr[i])
	}
	return res
}

// @lc code=end

func TestNextGreaterNodeInLinkedList(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,4,3,5]\n
// @lcpr case=end

*/
