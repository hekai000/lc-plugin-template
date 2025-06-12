/*
 * @lc app=leetcode.cn id=460 lang=golang
 * @lcpr version=30104
 *
 * [460] LFU 缓存
 */

package leetcode_solutions

import (
	"container/list"
	"testing"
)

// @lc code=start
type LFUCache struct {
	capacity int
	kvtable  map[int]int
	kftable  map[int]int
	fktable  map[int]*list.List
	kntable  map[int]*list.Element
	minfreq  int
	size     int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		kvtable:  make(map[int]int),
		kftable:  make(map[int]int),
		fktable:  make(map[int]*list.List),
		kntable:  make(map[int]*list.Element),
		minfreq:  0,
		size:     0,
	}
}

// 获取缓存中指定key的值
func (this *LFUCache) Get(key int) int {

	if _, ok := this.kvtable[key]; !ok {
		return -1
	}
	this.increaseFreq(key)

	return this.kvtable[key]
}

func (this *LFUCache) increaseFreq(key int) {
	//更新kf表
	freq := this.kftable[key]
	this.kftable[key] = freq + 1

	//更新fk表
	node := this.kntable[key]
	this.fktable[freq].Remove(node)
	if this.fktable[freq].Len() == 0 {
		delete(this.fktable, freq)
		if freq == this.minfreq {
			this.minfreq++
		}
	}
	if _, ok := this.fktable[freq+1]; !ok {
		this.fktable[freq+1] = list.New()
	}

	this.kntable[key] = this.fktable[freq+1].PushBack(key)

}
func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}
	if _, ok := this.kvtable[key]; ok {
		this.kvtable[key] = value
		this.increaseFreq(key)
		return
	}

	if this.size >= this.capacity {
		this.removeMinFreqKey()
	}

	this.kvtable[key] = value

	this.kftable[key] = 1

	if _, ok := this.fktable[1]; !ok {
		this.fktable[1] = list.New()
	}
	this.kntable[key] = this.fktable[1].PushBack(key)
	this.minfreq = 1
	this.size++

}

func (this *LFUCache) removeMinFreqKey() {
	keys := this.fktable[this.minfreq]
	node := keys.Front()
	deletedKey := node.Value.(int)
	keys.Remove(node)
	if keys.Len() == 0 {
		delete(this.fktable, this.minfreq)
	}
	delete(this.kvtable, deletedKey)
	delete(this.kftable, deletedKey)
	delete(this.kntable, deletedKey)
	this.size--
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func TestLfuCache(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]\n[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]\n
// @lcpr case=end

*/
