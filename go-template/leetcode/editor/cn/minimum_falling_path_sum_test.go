/*
 * @lc app=leetcode.cn id=931 lang=golang
 * @lcpr version=30104
 *
 * [931] 下降路径最小和
 */

package leetcode_solutions

import (
	"math"
	"testing"
)

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	//dp(i,j)定义为从matrix[0][...]下落到i,j位置的路径和
	n := len(matrix)
	res := math.MaxInt32
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = 66666
		}
	}
	for i := 0; i < n; i++ {
		res = minFPS2(res, dpFPS(matrix, n-1, i, memo))
	}
	return res
}

func dpFPS(matrix [][]int, i int, j int, memo [][]int) int {
	n := len(matrix)
	if i < 0 || j < 0 || i >= n || j >= n {
		return 99999
	}
	if i == 0 {
		return matrix[0][j]
	}
	if memo[i][j] != 66666 {
		return memo[i][j]
	}
	memo[i][j] = matrix[i][j] + minFPS3(dpFPS(matrix, i-1, j-1, memo), dpFPS(matrix, i-1, j, memo), dpFPS(matrix, i-1, j+1, memo))
	return memo[i][j]

}

func minFPS2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minFPS3(a, b, c int) int {
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

func TestMinimumFallingPathSum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[2,1,3],[6,5,4],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[-19,57],[-40,-5]]\n
// @lcpr case=end

*/
