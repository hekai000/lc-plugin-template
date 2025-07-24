/*
 * @lc app=leetcode.cn id=384 lang=golang
 * @lcpr version=30104
 *
 * [384] 打乱数组
 */

package leetcode_solutions

import (
	"math/rand"
	"testing"
	"time"
)

// @lc code=start
type Solution struct {
	nums []int
	rand *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	var cp = make([]int, n)
	copy(cp, this.nums)
	for i := 0; i < n; i++ {
		r := i + this.rand.Intn(n-i)
		cp[i], cp[r] = cp[r], cp[i]
	}
	return cp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

func TestShuffleAnArray(t *testing.T) {
	// your test code here

}
