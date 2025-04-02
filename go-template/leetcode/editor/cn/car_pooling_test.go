/*
 * @lc app=leetcode.cn id=1094 lang=golang
 * @lcpr version=30104
 *
 * [1094] 拼车
 */

package leetcode_solutions

import "testing"

// @lc code=start
type Difference2 struct {
	// 差分数组
	diff []int
}

// 输入一个初始数组，区间操作将在这个数组上进行
func NewDifference2(nums []int) *Difference2 {
	diff := make([]int, len(nums))
	// 根据初始数组构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference2{diff: diff}
}

// 给闭区间 [i, j] 增加 val（可以是负数）
func (d *Difference2) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// 返回结果数组
func (d *Difference2) Result() []int {
	res := make([]int, len(d.diff))
	// 根据差分数组构造结果数组
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	df := NewDifference2(nums)
	for _, v := range trips {
		val := v[0]
		i := v[1]
		j := v[2] - 1
		df.Increment(i, j, val)

	}
	res := df.Result()
	for _, v := range res {
		if v > capacity {
			return false
		}
	}
	return true
}

// @lc code=end

func TestCarPooling(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[2,1,5],[3,3,7]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,5],[3,3,7]]\n5\n
// @lcpr case=end

*/
