/*
 * @lc app=leetcode.cn id=1670 lang=golang
 * @lcpr version=30104
 *
 * [1670] 设计前中后队列
 */

package leetcode_solutions

import (
	"container/list"
	"testing"
)

// @lc code=start
type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

func ConstructorCC() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

func (this *FrontMiddleBackQueue) balance() {
	if this.right.Len() > this.left.Len()+1 {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.left.PushFront(val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if (this.left.Len()+this.right.Len())%2 == 0 {
		this.right.PushFront(val)
	} else {
		this.left.PushBack(val)
	}
	this.balance()
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.right.PushBack(val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	if this.left.Len()+this.right.Len() == 1 {
		return this.right.Remove(this.right.Front()).(int)
	}
	e := this.left.Remove(this.left.Front()).(int)
	this.balance()
	return e
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	var e int
	if (this.left.Len()+this.right.Len())%2 == 0 {
		e = this.left.Remove(this.left.Back()).(int)
	} else {
		// 如果有奇数个元素时，popMiddle 优先从右边删除
		e = this.right.Remove(this.right.Front()).(int)
	}
	this.balance()
	return e
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.left.Len()+this.right.Len() == 0 {
		return -1
	}
	e := this.right.Remove(this.right.Back()).(int)
	this.balance()
	return e
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
// @lc code=end

func TestDesignFrontMiddleBackQueue(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle",\n"popBack", "popFront"]\n[[], [1], [2], [3], [4], [], [], [], [], []]\n
// @lcpr case=end

*/
