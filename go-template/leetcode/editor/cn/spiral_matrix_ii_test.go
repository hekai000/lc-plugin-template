/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30103
 *
 * [59] 螺旋矩阵 II
 */

package leetcode_solutions

import "testing"

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for row := range matrix {
		matrix[row] = make([]int, n)
	}
	upperBound, lowerBound, leftBound, rightBound := 0, n-1, 0, n-1
	c := 1
	for c <= n*n {
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				matrix[upperBound][j] = c
				c++
			}
			upperBound++
		}

		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				matrix[i][rightBound] = c
				c++

			}
			rightBound--
		}

		if upperBound <= lowerBound {
			for j := rightBound; j >= leftBound; j-- {
				matrix[lowerBound][j] = c
				c++

			}
			lowerBound--
		}

		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				matrix[i][leftBound] = c
				c++
			}
			leftBound++
		}
	}
	return matrix
}

// @lc code=end

func TestSpiralMatrixIi(t *testing.T) {
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
