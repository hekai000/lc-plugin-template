/*
 * @lc app=leetcode.cn id=1260 lang=golang
 * @lcpr version=30103
 *
 * [1260] 二维网格迁移
 */

package leetcode_solutions

import "testing"

// @lc code=start

func shiftGrid(grid [][]int, k int) [][]int {

	m, n := len(grid), len(grid[0])
	k = k % (m * n)
	if k == 0 {
		return grid
	}
	reverseGrid(grid, m*n-k, m*n-1)
	reverseGrid(grid, 0, m*n-k-1)
	reverseGrid(grid, 0, m*n-1)

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = grid[i][j]
		}
	}
	return res

}

func reverseGrid(grid [][]int, i int, j int) {
	for i < j {
		temp := getElem(grid, i)
		setElem(grid, i, getElem(grid, j))
		setElem(grid, j, temp)
		i++
		j--
	}
}

func getElem(grid [][]int, index int) int {
	n := len(grid[0])
	i, j := index/n, index%n
	return grid[i][j]
}

func setElem(grid [][]int, index int, val int) {
	n := len(grid[0])
	i, j := index/n, index%n
	grid[i][j] = val

}
func shiftGrid2(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	temp := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp = append(temp, grid[i][j])
		}
	}
	k = k % (m * n)
	if k == 0 {
		return grid
	}
	temp = append(temp[(n*m)-k:], temp[:m*n-k]...)
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = temp[index]
			index++
		}
	}
	return grid
}

// @lc code=end

func TestShift2dGrid(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// your test code here
	shiftGrid(grid, 2)

}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n9\n
// @lcpr case=end

*/
