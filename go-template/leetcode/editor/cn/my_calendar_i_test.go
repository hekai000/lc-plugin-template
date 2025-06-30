/*
 * @lc app=leetcode.cn id=729 lang=golang
 * @lcpr version=30104
 *
 * [729] 我的日程安排表 I
 */

package leetcode_solutions

import (
	"container/list"
	"math"
	"testing"
)

// @lc code=start

type AllInOneSegmentTree struct {
	root         *SegmentNode
	defaultValue int
	typeOfMerge  string
}

type SegmentNode struct {
	l, r          int
	mergeValue    int
	left, right   *SegmentNode
	lazyAdd       int
	lazyAssign    int
	hasLazyAssign bool
}

func NewAllInOneSegmentTree(start, end, defaultValue int, t string) *AllInOneSegmentTree {
	if t != "sum" && t != "max" && t != "min" {
		panic("Invalid type, must be sum, max, or min")
	}
	rootMergeValue := computeRangeValue(start, end, defaultValue, t)
	root := &SegmentNode{
		l:             start,
		r:             end,
		mergeValue:    rootMergeValue,
		lazyAdd:       0,
		lazyAssign:    0,
		hasLazyAssign: false,
	}
	return &AllInOneSegmentTree{
		root:         root,
		defaultValue: defaultValue,
		typeOfMerge:  t,
	}
}

// 计算区间 [l, r] 赋值为 val 后的 mergeValue
func computeRangeValue(l, r, val int, t string) int {
	// 如果类型为求和，则返回区间长度乘以 val
	if t == "sum" {
		return (r - l + 1) * val
	} else {
		// 如果类型为求最大值或最小值，则返回 val
		return val
	}
}

func (tree *AllInOneSegmentTree) applyAddToValue(node *SegmentNode, delta int) int {
	if tree.typeOfMerge == "sum" {
		return node.mergeValue + (node.r-node.l+1)*delta
	} else {
		// 如果类型为求最大值或最小值，则返回当前值加上 delta
		return node.mergeValue + delta
	}
}

func (tree *AllInOneSegmentTree) merge(leftVal, rightVal int) int {
	if tree.typeOfMerge == "sum" {
		return leftVal + rightVal
	} else if tree.typeOfMerge == "max" {
		return int(math.Max(float64(leftVal), float64(rightVal)))
	} else if tree.typeOfMerge == "min" {
		return int(math.Min(float64(leftVal), float64(rightVal)))
	}
	panic("Invalid type")
}

func (tree *AllInOneSegmentTree) initChildrenIfNeeded(node *SegmentNode) {
	if node.l == node.r {
		return
	}
	mid := node.l + (node.r-node.l)/2
	if node.left == nil {
		leftMergeValue := computeRangeValue(node.l, mid, tree.defaultValue, tree.typeOfMerge)
		node.left = &SegmentNode{l: node.l, r: mid, mergeValue: leftMergeValue}
	}
	if node.right == nil {
		rightMergeValue := computeRangeValue(mid+1, node.r, tree.defaultValue, tree.typeOfMerge)
		node.right = &SegmentNode{l: mid + 1, r: node.r, mergeValue: rightMergeValue}
	}
}

func (tree *AllInOneSegmentTree) pushDown(node *SegmentNode) {
	// 如果存在赋值懒标记，优先下传赋值
	if node.hasLazyAssign {
		tree.applyAssign(node.left, node.lazyAssign)
		tree.applyAssign(node.right, node.lazyAssign)
		node.hasLazyAssign = false
		node.lazyAssign = 0
	}
	// 下传累加懒标记
	if node.lazyAdd != 0 {
		tree.applyAdd(node.left, node.lazyAdd)
		tree.applyAdd(node.right, node.lazyAdd)
		node.lazyAdd = 0
	}
}

func (tree *AllInOneSegmentTree) applyAssign(child *SegmentNode, val int) {
	child.hasLazyAssign = true
	child.lazyAssign = val
	// 清除累加懒标记
	child.lazyAdd = 0
	child.mergeValue = computeRangeValue(child.l, child.r, val, tree.typeOfMerge)
}

func (tree *AllInOneSegmentTree) applyAdd(child *SegmentNode, delta int) {
	if child.hasLazyAssign {
		// 如果子节点已有赋值懒标记，则直接更新该赋值
		child.lazyAssign += delta
		child.mergeValue = computeRangeValue(child.l, child.r, child.lazyAssign, tree.typeOfMerge)
	} else {
		// 如果子节点没有赋值懒标记，则更新累加懒标记
		child.lazyAdd += delta
		child.mergeValue = tree.applyAddToValue(child, delta)
	}
}

func (tree *AllInOneSegmentTree) update(index, val int) {
	// 直接复用区间赋值方法
	tree.RangeUpdate(index, index, val)
}

func (tree *AllInOneSegmentTree) RangeUpdate(qL, qR, val int) {
	tree._rangeUpdate(tree.root, qL, qR, val)
}

func (tree *AllInOneSegmentTree) _rangeUpdate(node *SegmentNode, qL, qR, val int) {
	if node.r < qL || node.l > qR {
		panic("Invalid update range")
	}
	// 当前节点完全包含于更新区间内
	if qL <= node.l && node.r <= qR {
		node.hasLazyAssign = true
		node.lazyAssign = val
		node.lazyAdd = 0
		node.mergeValue = computeRangeValue(node.l, node.r, val, tree.typeOfMerge)
		return
	}

	tree.initChildrenIfNeeded(node)
	tree.pushDown(node)

	mid := node.l + (node.r-node.l)/2
	if qR <= mid {
		tree._rangeUpdate(node.left, qL, qR, val)
	} else if qL > mid {
		tree._rangeUpdate(node.right, qL, qR, val)
	} else {
		tree._rangeUpdate(node.left, qL, mid, val)
		tree._rangeUpdate(node.right, mid+1, qR, val)
	}
	node.mergeValue = tree.merge(node.left.mergeValue, node.right.mergeValue)
}

func (tree *AllInOneSegmentTree) RangeAdd(qL, qR, delta int) {
	tree._rangeAdd(tree.root, qL, qR, delta)
}

func (tree *AllInOneSegmentTree) _rangeAdd(node *SegmentNode, qL, qR, delta int) {
	if node.r < qL || node.l > qR {
		panic("Invalid update range")
	}
	if qL <= node.l && node.r <= qR {
		if node.hasLazyAssign {
			// 若已有赋值懒标记，则更新赋值值
			node.lazyAssign += delta
			node.mergeValue = computeRangeValue(node.l, node.r, node.lazyAssign, tree.typeOfMerge)
		} else {
			node.lazyAdd += delta
			node.mergeValue = tree.applyAddToValue(node, delta)
		}
		return
	}
	tree.initChildrenIfNeeded(node)
	tree.pushDown(node)

	mid := node.l + (node.r-node.l)/2
	if qL <= mid {
		tree._rangeAdd(node.left, qL, qR, delta)
	}
	if qR > mid {
		tree._rangeAdd(node.right, qL, qR, delta)
	}
	node.mergeValue = tree.merge(node.left.mergeValue, node.right.mergeValue)
}

func (tree *AllInOneSegmentTree) Query(qL, qR int) int {
	return tree._query(tree.root, qL, qR)
}

func (tree *AllInOneSegmentTree) _query(node *SegmentNode, qL, qR int) int {
	if node.r < qL || node.l > qR {
		panic("Invalid query range")
	}
	if qL <= node.l && node.r <= qR {
		return node.mergeValue
	}

	tree.initChildrenIfNeeded(node)
	tree.pushDown(node)

	mid := node.l + (node.r-node.l)/2
	if qR <= mid {
		return tree._query(node.left, qL, qR)
	} else if qL > mid {
		return tree._query(node.right, qL, qR)
	} else {
		leftResult := tree._query(node.left, qL, mid)
		rightResult := tree._query(node.right, mid+1, qR)
		return tree.merge(leftResult, rightResult)
	}
}

// type MyCalendar struct {
// 	calendar *AllInOneSegmentTree
// }

// func Constructor() MyCalendar {
// 	myCalendar := NewAllInOneSegmentTree(0, 1000000000, 0, "max")
// 	return MyCalendar{
// 		calendar: myCalendar,
// 	}
// }

// func (this *MyCalendar) Book(startTime int, endTime int) bool {
// 	if this.calendar.Query(startTime, endTime-1) > 0 {
// 		return false
// 	}
// 	this.calendar.RangeAdd(startTime, endTime-1, 1)
// 	return true
// }

type MyCalendar struct {
	calendar *list.List
}

func Constructor55() MyCalendar {

	return MyCalendar{
		calendar: list.New(),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for e := this.calendar.Front(); e != nil; e = e.Next() {
		event := e.Value.([2]int)
		if event[0] < endTime && startTime < event[1] {
			return false
		}
	}
	this.calendar.PushBack([2]int{startTime, endTime})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
// @lc code=end

func TestMyCalendarI(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["MyCalendar", "book", "book", "book"]\n[[], [10, 20], [15, 25], [20, 30]]\n
// @lcpr case=end

*/
