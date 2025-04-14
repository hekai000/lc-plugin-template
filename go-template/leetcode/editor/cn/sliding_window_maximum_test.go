/*
 * @lc app=leetcode.cn id=239 lang=golang
 * @lcpr version=30104
 *
 * [239] 滑动窗口最大值
 */

package leetcode_solutions

import (
	"container/list"
	"testing"
)

// @lc code=start

type MontonicQueue struct {
	maxq list.List
}

func (mq *MontonicQueue) push(n int) {
	for mq.maxq.Len() > 0 && n > mq.maxq.Back().Value.(int) {
		mq.maxq.Remove(mq.maxq.Back())
	}
	mq.maxq.PushBack(n)
}

func (mq *MontonicQueue) pop(n int) {
	if n == mq.maxq.Front().Value.(int) {
		mq.maxq.Remove(mq.maxq.Front())
	}
}

func (mq *MontonicQueue) max() int {
	return mq.maxq.Front().Value.(int)
}
func maxSlidingWindow(nums []int, k int) []int {
	window := MontonicQueue{maxq: list.List{}}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
		} else {
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return res
}

// @lc code=end

func TestSlidingWindowMaximum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,3,-1,-3,5,3,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
