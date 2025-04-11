/*
 * @lc app=leetcode.cn id=641 lang=golang
 * @lcpr version=30104
 *
 * [641] 设计循环双端队列
 */

package leetcode_solutions

import (
	"errors"
	"testing"
)

// @lc code=start
const INIT_CAP = 2

type MyArrayDeque struct {
	size        int
	data        []int
	first, last int
}

func NewMyArrayDeque(initCap int) *MyArrayDeque {
	return &MyArrayDeque{
		size:  0,
		data:  make([]int, initCap),
		first: 0,
		last:  0,
	}
}
func NewMyArrayDequeDefault() *MyArrayDeque {
	return NewMyArrayDeque(INIT_CAP)
}

func (d *MyArrayDeque) Size() int {
	return d.size
}

func (d *MyArrayDeque) resize(newCap int) {
	temp := make([]int, newCap)
	for i := 0; i < d.size; i++ {
		temp[i] = d.data[(d.first+i)%len(d.data)]
	}
	d.first = 0
	d.last = d.size
	d.data = temp
}
func (d *MyArrayDeque) AddFirst(value int) {
	if d.size == len(d.data) {
		d.resize(d.size * 2)
	}
	if d.first == 0 {
		d.first = len(d.data) - 1
	} else {
		d.first--
	}
	d.data[d.first] = value
	d.size++
}

func (d *MyArrayDeque) AddLast(value int) {
	if d.size == len(d.data) {
		d.resize(d.size * 2)
	}
	d.data[d.last] = value
	d.last++
	if d.last == len(d.data) {
		d.last = 0
	}
	d.size++
}

func (d *MyArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyArrayDeque) RemoveLast() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("NoSuchElem")
	}
	if d.size == len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}
	if d.last == 0 {
		d.last = len(d.data) - 1
	} else {
		d.last--
	}
	oldVal := d.data[d.last]
	d.data[d.last] = 0
	d.size--
	return oldVal, nil
}

func (d *MyArrayDeque) RemoveFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("NoSuchElem")
	}
	if d.size == len(d.data)/4 {
		d.resize(len(d.data) / 2)
	}
	oldVal := d.data[d.first]
	d.data[d.first] = 0
	d.first++
	if d.first == len(d.data) {
		d.first = 0
	}
	d.size--
	return oldVal, nil
}

func (d *MyArrayDeque) GetFirst() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("NoSuchElem")
	}
	return d.data[d.first], nil
}

func (d *MyArrayDeque) GetLast() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("NoSuchElem")
	}
	if d.last == 0 {
		return d.data[len(d.data)-1], nil
	}
	return d.data[d.last-1], nil
}

type MyCircularDeque struct {
	cap  int
	list *MyArrayDeque
}

func Constructor44(k int) MyCircularDeque {
	return MyCircularDeque{
		cap:  k,
		list: NewMyArrayDequeDefault(),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.list.Size() == this.cap {
		return false
	}
	this.list.AddFirst(value)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.list.Size() == this.cap {
		return false
	}
	this.list.AddLast(value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.list.IsEmpty() {
		return false
	}
	_, _ = this.list.RemoveFront()
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.list.IsEmpty() {
		return false
	}
	_, _ = this.list.RemoveLast()
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.list.IsEmpty() {
		return -1
	}
	val, _ := this.list.GetFirst()
	return val
}

func (this *MyCircularDeque) GetRear() int {
	if this.list.IsEmpty() {
		return -1
	}
	val, _ := this.list.GetLast()
	return val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.list.IsEmpty()
}

func (this *MyCircularDeque) IsFull() bool {
	return this.list.Size() == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

func TestDesignCircularDeque(t *testing.T) {
	// your test code here

}
