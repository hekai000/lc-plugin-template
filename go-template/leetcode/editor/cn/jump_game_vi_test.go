/*
 * @lc app=leetcode.cn id=1696 lang=golang
 * @lcpr version=30104
 *
 * [1696] 跳跃游戏 VI
 */

package leetcode_solutions

import (
	"testing"
)

// @lc code=start
func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n)
	window := NewMyMonotonicQueue4()

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = -1 << 32
	}
	window.push(dp[0])
	for i := 1; i < n; i++ {
		dp[i] = window.max() + nums[i]
		if window.size() == k {
			window.pop()
		}
		window.push(dp[i])
	}
	return dp[n-1]

}

type MyMonotonicQueue4 struct {
	maxq []int
	minq []int
	q    []int
}

func NewMyMonotonicQueue4() *MyMonotonicQueue4 {
	return &MyMonotonicQueue4{
		maxq: make([]int, 0),
		minq: make([]int, 0),
		q:    make([]int, 0),
	}
}

func (mq *MyMonotonicQueue4) push(e int) {
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

func (mq *MyMonotonicQueue4) pop() int {
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

func (mq *MyMonotonicQueue4) max() int {
	return mq.maxq[0]
}

func (mq *MyMonotonicQueue4) min() int {
	return mq.minq[0]
}

func (mq *MyMonotonicQueue4) size() int {
	return len(mq.q)
}

func (mq *MyMonotonicQueue4) isEmpty() bool {
	return len(mq.q) == 0
}

// @lc code=end

func TestJumpGameVi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,-1,-2,4,-7,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [10,-5,-2,4,0,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,-5,-20,4,-1,3,-6,-3]\n2\n
// @lcpr case=end

*/
