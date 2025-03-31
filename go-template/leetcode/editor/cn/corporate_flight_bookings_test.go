/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=30104
 *
 * [1109] 航班预订统计
 */

package leetcode_solutions

import "testing"

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n+1)

	for _, v := range bookings {
		for j := v[0]; j <= v[1]; j++ {
			res[j] += v[2]
		}
	}
	return res[1:]

}

// @lc code=end

func TestCorporateFlightBookings(t *testing.T) {
	// your test code here
	a := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	corpFlightBookings(a, 5)

}

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/
