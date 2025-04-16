/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=30104
 *
 * [918] 环形子数组的最大和
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	preSum := make([]int, 2*n+1)
	preSum[0] = 0

	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[(i-1)%n]
	}
	maxSum := math.MinInt32
	window := NewMontonicQueue3()
	window.push(0)

	for i := 1; i < len(preSum); i++ {
		maxSum = max(maxSum, preSum[i]-window.min())
		if window.size() == n {
			window.pop()
		}
		window.push(preSum[i])
	}
	return maxSum
}

type MyMonotonicQueue3 struct {
	maxq []int
	minq []int
	q    []int
}

func NewMontonicQueue3() *MyMonotonicQueue3 {
	return &MyMonotonicQueue3{
		maxq: make([]int, 0),
		minq: make([]int, 0),
		q:    make([]int, 0),
	}
}

func (mq *MyMonotonicQueue3) push(e int) {
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

func (mq *MyMonotonicQueue3) pop() int {
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

func (mq *MyMonotonicQueue3) max() int {
	return mq.maxq[0]
}

func (mq *MyMonotonicQueue3) min() int {
	return mq.minq[0]
}

func (mq *MyMonotonicQueue3) size() int {
	return len(mq.q)
}

func (mq *MyMonotonicQueue3) isEmpty() bool {
	return len(mq.q) == 0
}

// @lc code=end

func TestMaximumSumCircularSubarray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/
