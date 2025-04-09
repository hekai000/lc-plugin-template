/*
 * @lc app=leetcode.cn id=155 lang=golang
 * @lcpr version=30104
 *
 * [155] 最小栈
 */

package leetcode_solutions

import "testing"

// @lc code=start
type MinStack struct {
	stk    []int
	minStk []int
}

func Constructor7() MinStack {
	return MinStack{
		stk:    make([]int, 0),
		minStk: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	if len(this.minStk) == 0 || val <= this.minStk[len(this.minStk)-1] {
		this.minStk = append(this.minStk, val)
	} else {
		this.minStk = append(this.minStk, this.minStk[len(this.minStk)-1])
	}
}

func (this *MinStack) Pop() {
	this.stk = this.stk[:len(this.stk)-1]
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

func TestMinStack(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]\n
// @lcpr case=end

*/
