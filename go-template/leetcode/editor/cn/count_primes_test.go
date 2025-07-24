/*
 * @lc app=leetcode.cn id=204 lang=golang
 * @lcpr version=30104
 *
 * [204] 计数质数
 */

package leetcode_solutions

import "testing"

// @lc code=start
func countPrimes(n int) int {
	count := 0
	isPrimes := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrimes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrimes[i] {
			for j := i * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}

	for i := 2; i < n; i++ {
		if isPrimes[i] {
			count++
		}
	}
	return count
}

// @lc code=end

func TestCountPrimes(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
