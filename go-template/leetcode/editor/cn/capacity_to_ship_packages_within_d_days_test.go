/*
 * @lc app=leetcode.cn id=1011 lang=golang
 * @lcpr version=30104
 *
 * [1011] 在 D 天内送达包裹的能力
 */

package leetcode_solutions

import "testing"

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	//运载能力k，f(k)表示在运载能力为k时，需要f(k)天才能运完。
	left, right := 0, 1
	for _, w := range weights {
		if left < w {
			left = w
		}
		right += w
	}
	for left <= right {
		mid := left + (right-left)/2
		if fk(weights, mid) <= days {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func fk(weights []int, k int) int {
	days := 0
	for i := 0; i < len(weights); {
		cap := k
		for i < len(weights) {
			if cap < weights[i] {
				break
			} else {
				cap -= weights[i]
			}
			i++
		}
		days++

	}
	return days
}

// @lc code=end

func TestCapacityToShipPackagesWithinDDays(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5,6,7,8,9,10]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,2,2,4,1,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,1]\n4\n
// @lcpr case=end

*/
