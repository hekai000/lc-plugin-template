/*
 * @lc app=leetcode.cn id=1438 lang=golang
 * @lcpr version=30104
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestSubarray(nums []int, limit int) int {
	left, right := 0, 0
	windowSize, res := 0, 0
	window := NewMontonicQueue()
	for right < len(nums) {
		window.push(nums[right])
		windowSize++
		right++

		for window.max()-window.min() > limit {
			window.pop()
			left++
			windowSize--
		}
		if windowSize > res {
			res = windowSize
		}

	}
	return res
}

type MyMonotonicQueue struct {
	maxq []int
	minq []int
	q    []int
}

func NewMontonicQueue() *MyMonotonicQueue {
	return &MyMonotonicQueue{
		maxq: make([]int, 0),
		minq: make([]int, 0),
		q:    make([]int, 0),
	}
}

func (mq *MyMonotonicQueue) push(e int) {
	mq.q = append(mq.q, e)

	for len(mq.maxq) > 0 && e > mq.maxq[len(mq.maxq)-1] {
		mq.maxq = mq.maxq[:len(mq.maxq)-1]
	}
	mq.maxq = append(mq.maxq, e)
	for len(mq.minq) > 0 && e < mq.minq[len(mq.minq)-1] {
		mq.minq = mq.minq[:len(mq.minq)-1]
	}
	mq.minq = append(mq.minq, e)

}

func (mq *MyMonotonicQueue) pop() int {
	deleteVal := mq.q[0]
	mq.q = mq.q[1:]

	if deleteVal == mq.maxq[0] {
		mq.maxq = mq.maxq[1:]
	}

	if deleteVal == mq.minq[0] {
		mq.minq = mq.minq[1:]
	}
	return deleteVal
}

func (mq *MyMonotonicQueue) max() int {
	return mq.maxq[0]
}

func (mq *MyMonotonicQueue) min() int {
	return mq.minq[0]
}

func (mq *MyMonotonicQueue) size() int {
	return len(mq.q)
}

func (mq *MyMonotonicQueue) isEmpty() bool {
	return len(mq.q) == 0
}

// @lc code=end

func TestLongestContinuousSubarrayWithAbsoluteDiffLessThanOrEqualToLimit(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [8,2,4,7]\n4\n
// @lcpr case=end

// @lcpr case=start
// [10,1,2,4,7,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [4,2,2,2,4,4,2,2]\n0\n
// @lcpr case=end

*/
