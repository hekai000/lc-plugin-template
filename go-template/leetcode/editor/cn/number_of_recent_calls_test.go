/*
 * @lc app=leetcode.cn id=933 lang=golang
 * @lcpr version=30104
 *
 * [933] 最近的请求次数
 */

package leetcode_solutions

import "testing"

// @lc code=start
type RecentCounter struct {
	q []int
}

func Constructor12() RecentCounter {
	return RecentCounter{q: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

func TestNumberOfRecentCalls(t *testing.T) {
	// your test code here
	obj := Constructor12()
	obj.Ping(1)
}

/*
// @lcpr case=start
// ["RecentCounter", "ping", "ping", "ping", "ping"]\n[[], [1], [100], [3001], [3002]]\n
// @lcpr case=end

*/
