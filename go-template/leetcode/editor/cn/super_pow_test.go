/*
 * @lc app=leetcode.cn id=372 lang=golang
 * @lcpr version=30104
 *
 * [372] 超级次方
 */

package leetcode_solutions

import "testing"

// @lc code=start
func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]

	part1 := mypow22(a, last)
	part2 := mypow22(superPow(a, b), 10)

	return (part1 * part2) % 1337

}

func mypow22(a, k int) int {
	a %= 1337
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= 1337
	}
	return res
}

// @lc code=end

func TestSuperPow(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 2\n[3]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,0]\n
// @lcpr case=end

// @lcpr case=start
// 1\n[4,3,3,8,5,2]\n
// @lcpr case=end

// @lcpr case=start
// 2147483647\n[2,0,0]\n
// @lcpr case=end

*/
