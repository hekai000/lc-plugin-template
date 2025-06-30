/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=30104
 *
 * [146] LRU 缓存
 */

package leetcode_solutions

import (
	"container/list"
	"testing"
)

// @lc code=start
type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[int]*list.Element
}

type pair struct {
	key   int
	value int
}

func Constructor33(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.list.MoveToFront(ele)
		return ele.Value.(*pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		this.list.MoveToFront(ele)
		ele.Value.(*pair).value = value
		return
	}

	if this.list.Len() >= this.capacity {
		last := this.list.Back()
		delete(this.cache, last.Value.(*pair).key)
		this.list.Remove(last)
	}
	ele := this.list.PushFront(&pair{key, value})
	this.cache[key] = ele

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func TestLruCache(t *testing.T) {
	// your test code here

}
