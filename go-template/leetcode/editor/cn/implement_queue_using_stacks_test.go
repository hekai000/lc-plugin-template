/*
 * @lc app=leetcode.cn id=232 lang=golang
 * @lcpr version=30104
 *
 * [232] 用栈实现队列
 */

package leetcode_solutions

import "testing"

// @lc code=start
type MyQueue struct {
	s1, s2 []int
}

func Constructor5() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Peek() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	return this.s2[len(this.s2)-1]
}

func (this *MyQueue) Pop() int {
	this.Peek()
	retVal := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return retVal
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

func TestImplementQueueUsingStacks(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["MyQueue", "push", "push", "peek", "pop", "empty"]\n[[], [1], [2], [], [], []]\n
// @lcpr case=end

*/
