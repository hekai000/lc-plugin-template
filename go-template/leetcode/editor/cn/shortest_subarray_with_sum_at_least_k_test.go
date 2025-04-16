/*
 * @lc app=leetcode.cn id=862 lang=golang
 * @lcpr version=30104
 *
 * [862] 和至少为 K 的最短子数组
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func shortestSubarray(nums []int, k int) int {
	window := NewMontonicQueue2()
	n := len(nums)
	preSum := make([]int64, n+1)
	preSum[0] = 0
	minLen := math.MaxInt64
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + int64(nums[i-1])
	}

	left, right := 0, 0
	for right < len(preSum) {
		window.push(preSum[right])
		right++

		for right < len(preSum) && !window.isEmpty() && preSum[right]-window.min() >= int64(k) {
			minLen = int(min(minLen, right-left))
			window.pop()
			left++
		}

	}
	if minLen == math.MaxInt64 {
		return -1
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type MyMonotonicQueue2 struct {
	maxq []int64
	minq []int64
	q    []int64
}

func NewMontonicQueue2() *MyMonotonicQueue2 {
	return &MyMonotonicQueue2{
		maxq: make([]int64, 0),
		minq: make([]int64, 0),
		q:    make([]int64, 0),
	}
}

func (mq *MyMonotonicQueue2) push(e int64) {
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

func (mq *MyMonotonicQueue2) pop() int64 {
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

func (mq *MyMonotonicQueue2) max() int64 {
	return mq.maxq[0]
}

func (mq *MyMonotonicQueue2) min() int64 {
	return mq.minq[0]
}

func (mq *MyMonotonicQueue2) size() int {
	return len(mq.q)
}

func (mq *MyMonotonicQueue2) isEmpty() bool {
	return len(mq.q) == 0
}

// @lc code=end

func TestShortestSubarrayWithSumAtLeastK(t *testing.T) {
	// your test code here
	nums := []int{1}
	k := 1
	shortestSubarray(nums, k)

}

/*
// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n4\n
// @lcpr case=end

// @lcpr case=start
// [2,-1,2]\n3\n
// @lcpr case=end

*/
