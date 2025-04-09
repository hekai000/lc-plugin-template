/*
 * @lc app=leetcode.cn id=225 lang=golang
 * @lcpr version=30104
 *
 * [225] 用队列实现栈
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
type MyStack struct {
	q        []int
	top_elem int
}

func Constructor4() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	this.top_elem = x
}

func (this *MyStack) Pop() int {
	size := len(this.q)
	for size > 2 {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
		size--
	}
	this.top_elem = this.q[0]
	this.q = append(this.q, this.q[0])
	this.q = this.q[1:]

	result := this.q[0]
	this.q = this.q[1:]
	return result
}

func (this *MyStack) Top() int {
	return this.top_elem
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

func TestImplementStackUsingQueues(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["MyStack", "push", "push", "top", "pop", "empty"]\n[[], [1], [2], [], [], []]\n
// @lcpr case=end

*/
