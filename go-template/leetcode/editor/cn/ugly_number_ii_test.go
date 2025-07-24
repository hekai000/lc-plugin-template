/*
 * @lc app=leetcode.cn id=264 lang=golang
 * @lcpr version=30100
 *
 * [264] 丑数 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1
	product2, product3, product5 := 1, 1, 1
	p := 1
	ugly := make([]int, n+1)
	for p <= n {
		min := minugly(product2, product3, product5)
		ugly[p] = min
		p++
		if min == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if min == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}

		if min == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}

	}
	return ugly[n]
}

func minugly(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

// @lc code=end

func TestUglyNumberIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
