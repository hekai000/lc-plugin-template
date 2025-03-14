/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30103
 *
 * [54] 螺旋矩阵
 */

package leetcode_solutions

import "testing"

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	m, n := len(matrix), len(matrix[0])
	upperBound, lowerBound, leftBound, rightBound := 0, m-1, 0, n-1
	for len(res) < m*n {
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upperBound][j])
			}
			upperBound++
		}

		if leftBound <= rightBound {
			for i := upperBound; i <= lowerBound; i++ {
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}

		if upperBound <= lowerBound {
			for j := rightBound; j >= leftBound; j-- {
				res = append(res, matrix[lowerBound][j])
			}
			lowerBound--
		}

		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}
	return res
}

// @lc code=end

func TestSpiralMatrix(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
