/*
 * @lc app=leetcode.cn id=96 lang=golang
 * @lcpr version=30104
 *
 * [96] 不同的二叉搜索树
 */

package leetcode_solutions

import "testing"

// @lc code=start
func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		memo[i] = make([]int, n+1)
	}
	var count func(low, high int) int
	count = func(low, high int) int {
		if low > high {
			return 1
		}
		if memo[low][high] != 0 {
			return memo[low][high]
		}
		res := 0
		for mid := low; mid <= high; mid++ {
			left := count(low, mid-1)
			right := count(mid+1, high)
			res += left * right
		}
		memo[low][high] = res
		return res
	}
	return count(1, n)
}

// @lc code=end

func TestUniqueBinarySearchTrees(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
