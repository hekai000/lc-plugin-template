/*
 * @lc app=leetcode.cn id=304 lang=golang
 * @lcpr version=30104
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

package leetcode_solutions

import "testing"

// @lc code=start
type NumMatrix struct {
	preSumMatrix [][]int
}

func Constructor21(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	preSumM := make([][]int, m+1)
	for i, _ := range preSumM {
		preSumM[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSumM[i][j] = preSumM[i-1][j] + preSumM[i][j-1] + matrix[i-1][j-1] - preSumM[i-1][j-1]
		}
	}
	return NumMatrix{preSumMatrix: preSumM}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSumMatrix[row2+1][col2+1] - this.preSumMatrix[row2+1][col1] - this.preSumMatrix[row1][col2+1] + this.preSumMatrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

func TestRangeSumQuery2dImmutable(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]\n[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]\n
// @lcpr case=end

*/
