/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=30104
 *
 * [45] 跳跃游戏 II
 */

package leetcode_solutions

import "testing"

// @lc code=start

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	jumps := 0
	farthest, end := 0, 0
	for i := 0; i < n-1; i++ {
		if nums[i]+i > farthest {
			farthest = nums[i] + i
		}
		if i == end {
			jumps++
			end = farthest
			if farthest >= n-1 {
				return jumps
			}
		}
	}
	return -1
}
func jump2(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = n
	}
	return dpjump(nums, 0, memo)
}

func dpjump(nums []int, i int, memo []int) int {
	n := len(nums)
	if i >= n-1 {
		return 0
	}
	if memo[i] != n {
		return memo[i]
	}

	steps := nums[i]
	for j := 1; j <= steps; j++ {
		subProblem := dpjump(nums, i+j, memo)
		memo[i] = minjump(memo[i], subProblem+1)
	}
	return memo[i]
}

func minjump(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func TestJumpGameIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/
