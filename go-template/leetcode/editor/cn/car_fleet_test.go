/*
 * @lc app=leetcode.cn id=853 lang=golang
 * @lcpr version=30104
 *
 * [853] 车队
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([][2]int, n)
	for i := 0; i < n; i++ {
		cars[i][0] = position[i]
		cars[i][1] = speed[i]
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	time := make([]float64, n)
	for i := 0; i < n; i++ {
		car := cars[i]
		time[i] = float64(target-car[0]) / float64(car[1])
	}

	res := 0

	maxTime := 0.0
	for i := n - 1; i >= 0; i-- {
		if time[i] > maxTime {
			maxTime = time[i]
			res++
		}
	}
	return res

}

// @lc code=end

func TestCarFleet(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 12\n[10,8,0,5,3]\n[2,4,1,1,3]\n
// @lcpr case=end

// @lcpr case=start
// 10\n[3]\n[3]\n
// @lcpr case=end

// @lcpr case=start
// 100\n[0,2,4]\n[4,2,1]\n
// @lcpr case=end

*/
