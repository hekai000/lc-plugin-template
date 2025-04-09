/*
 * @lc app=leetcode.cn id=895 lang=golang
 * @lcpr version=30104
 *
 * [895] 最大频率栈
 */

package leetcode_solutions

import "testing"

// @lc code=start
type FreqStack struct {
	maxFreq   int
	FreqVal   map[int][]int
	ValToFreq map[int]int
}

func Constructor9() FreqStack {
	return FreqStack{
		maxFreq:   0,
		FreqVal:   make(map[int][]int),
		ValToFreq: make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	if f, ok := this.ValToFreq[val]; ok {
		this.FreqVal[f+1] = append(this.FreqVal[f+1], val)
		this.ValToFreq[val] = f + 1
	} else {
		this.FreqVal[1] = append(this.FreqVal[1], val)
		this.ValToFreq[val] = 1
	}
	if this.ValToFreq[val] > this.maxFreq {
		this.maxFreq = this.ValToFreq[val]
	}

}

func (this *FreqStack) Pop() int {
	res := this.FreqVal[this.maxFreq][len(this.FreqVal[this.maxFreq])-1]
	this.FreqVal[this.maxFreq] = this.FreqVal[this.maxFreq][:len(this.FreqVal[this.maxFreq])-1]
	this.ValToFreq[res]--
	if len(this.FreqVal[this.maxFreq]) == 0 {
		delete(this.FreqVal, this.maxFreq)
		this.maxFreq--
	}
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

func TestMaximumFrequencyStack(t *testing.T) {
	// your test code here
	m := Constructor9()
	m.Push(5)
	m.Push(7)
	m.Push(5)
	m.Push(7)
	m.Push(4)
	m.Push(5)
	m.Pop()
	m.Pop()
	m.Pop()
	m.Pop()

}

/*
// @lcpr case=start
// ["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],\n[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]\n
// @lcpr case=end

*/
