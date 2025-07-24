/*
 * @lc app=leetcode.cn id=398 lang=golang
 * @lcpr version=30104
 *
 * [398] 随机数索引
 */

package leetcode_solutions

import (
	"math/rand"
	"testing"
	"time"
)

// @lc code=start
type Solution struct {
	nums      []int
	generator *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:      nums,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Pick(target int) int {
	res := 0
	var i int
	cnt := 0
	for i = 0; i < len(this.nums); i++ {
		if target == this.nums[i] {
			cnt++
			if this.generator.Intn(cnt) == 0 {
				res = i
			}
		}
	}
	return res

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

func TestRandomPickIndex(t *testing.T) {
	// your test code here

}
