/*
 * @lc app=leetcode.cn id=875 lang=golang
 * @lcpr version=30104
 *
 * [875] 爱吃香蕉的珂珂
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	left, right := 1, piles[len(piles)-1]

	for left <= right {
		mid := left + (right-left)/2
		if f(piles, mid) == h {
			right = mid - 1
		} else if f(piles, mid) < h {
			right = mid - 1
		} else if f(piles, mid) > h {
			left = mid + 1
		}

	}

	return left

}

func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
}

// @lc code=end

func TestKokoEatingBananas(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,6,7,11]\n8\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n5\n
// @lcpr case=end

// @lcpr case=start
// [30,11,23,4,20]\n6\n
// @lcpr case=end

*/
