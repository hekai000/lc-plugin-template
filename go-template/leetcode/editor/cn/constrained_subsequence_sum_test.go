/*
 * @lc app=leetcode.cn id=1425 lang=golang
 * @lcpr version=30104
 *
 * [1425] 带限制的子序列和
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	maxVal := math.MinInt32
	window := NewMontonicQueue5()
	dp := make([]int, n)
	dp[0] = nums[0]
	window.push(dp[0])
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], window.max()+nums[i])
		if window.size() == k {
			window.pop()
		}
		window.push(dp[i])
	}
	for i := 0; i < n; i++ {
		maxVal = max(maxVal, dp[i])
	}
	return maxVal
}

type MyMonotonicQueue5 struct {
	maxq []int
	minq []int
	q    []int
}

func NewMontonicQueue5() *MyMonotonicQueue5 {
	return &MyMonotonicQueue5{
		maxq: make([]int, 0),
		minq: make([]int, 0),
		q:    make([]int, 0),
	}
}

func (mq *MyMonotonicQueue5) push(e int) {
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

func (mq *MyMonotonicQueue5) pop() int {
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

func (mq *MyMonotonicQueue5) max() int {
	return mq.maxq[0]
}

func (mq *MyMonotonicQueue5) min() int {
	return mq.minq[0]
}

func (mq *MyMonotonicQueue5) size() int {
	return len(mq.q)
}

func (mq *MyMonotonicQueue5) isEmpty() bool {
	return len(mq.q) == 0
}

// @lc code=end

func TestConstrainedSubsequenceSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [10,2,-10,5,20]\n2\n
// @lcpr case=end

// @lcpr case=start
// [-1,-2,-3]\n1\n
// @lcpr case=end

// @lcpr case=start
// [10,-2,-10,-5,20]\n2\n
// @lcpr case=end

*/
