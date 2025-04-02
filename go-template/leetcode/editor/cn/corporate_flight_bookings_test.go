/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=30104
 *
 * [1109] 航班预订统计
 */

package leetcode_solutions

import "testing"

// @lc code=start
<<<<<<< HEAD

// 差分数组工具类
type Difference struct {
	// 差分数组
	diff []int
}

// 输入一个初始数组，区间操作将在这个数组上进行
func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))
	// 根据初始数组构造差分数组
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

// 给闭区间 [i, j] 增加 val（可以是负数）
func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// 返回结果数组
func (d *Difference) Result() []int {
	res := make([]int, len(d.diff))
	// 根据差分数组构造结果数组
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	df := NewDifference(nums)

	for _, booking := range bookings {
		i := booking[0] - 1
		j := booking[1] - 1
		val := booking[2]
		df.Increment(i, j, val)
	}
	return df.Result()
=======
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n+1)

	for _, v := range bookings {
		for j := v[0]; j <= v[1]; j++ {
			res[j] += v[2]
		}
	}
	return res[1:]

>>>>>>> 801e320f71632f83aadeb4f4703e3f79bcd61e82
}

// @lc code=end

func TestCorporateFlightBookings(t *testing.T) {
	// your test code here
<<<<<<< HEAD
=======
	a := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	corpFlightBookings(a, 5)
>>>>>>> 801e320f71632f83aadeb4f4703e3f79bcd61e82

}

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/
