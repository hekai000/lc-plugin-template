/*
 * @lc app=leetcode.cn id=622 lang=golang
 * @lcpr version=30104
 *
 * [622] 设计循环队列
 */

package leetcode_solutions

import "testing"

// @lc code=start

type node struct {
	val  int
	next *node
}
type MyCircularQueue struct {
	head     *node
	size     int
	capacity int
}

func Constructor11(k int) MyCircularQueue {
	return MyCircularQueue{
		head:     &node{},
		size:     0,
		capacity: k,
	}

}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.capacity == this.size {
		return false
	}
	p := this.head
	for p != nil && p.next != nil {
		p = p.next
	}
	p.next = &node{val: value}
	this.size += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	this.head.next = this.head.next.next
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.head.next.val
}

func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	p := this.head
	for p != nil && p.next != nil {
		p = p.next
	}
	return p.val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

func TestDesignCircularQueue(t *testing.T) {
	// your test code here

}
