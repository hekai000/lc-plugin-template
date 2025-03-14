/*
 * @lc app=leetcode.cn id=1329 lang=golang
 * @lcpr version=30103
 *
 * [1329] 将矩阵按对角线排序
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start

// 思路其实是差不多，我的思路是拿出每个对角线元素组成数组，然后排序，然后再把排序后的元素塞到对应对角线中去。
// labuladong的思路是，通过遍历二维mat将元素分类到一个map中，由于同一条对角线的i，j差值是一样的，所以可以以
// i-j为key，然后对每条对角线元素进行排序；这里排序注意是降序排列，然后在遍历时，从切片后面拿小的元素。
func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	diagonalMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagId := i - j
			diagonalMap[diagId] = append(diagonalMap[diagId], mat[i][j])
		}
	}

	for _, row := range diagonalMap {
		sort.Slice(row, func(i int, j int) bool {
			return row[i] > row[j]
		})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagId := i - j
			mat[i][j] = diagonalMap[diagId][len(diagonalMap[diagId])-1]
			diagonalMap[diagId] = diagonalMap[diagId][:(len(diagonalMap[diagId]) - 1)]
		}
	}
	return mat
}
func diagonalSort2(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	mat1 := make([][]int, m)
	for row := range mat1 {
		mat1[row] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			for j := 0; j <= n-1; j++ {
				v := getdiagE(mat, 0, j)
				setdiagE(mat1, i, j, v)
			}
		} else {
			v1 := getdiagE(mat, i, 0)
			setdiagE(mat1, i, 0, v1)
		}
	}
	return mat1
}

func getdiagE(mat [][]int, r int, c int) []int {
	m, n := len(mat), len(mat[0])
	res := []int{}
	i, j := r, c
	for ; i < m && j < n; i, j = i+1, j+1 {
		res = append(res, mat[i][j])
	}
	return res
}

// setidiagE 函数用于设置矩阵 mat 的对角线元素
// 参数 mat 是一个二维整数数组，表示矩阵
// 参数 r 是矩阵的行数
// 参数 c 是矩阵的列数
// 参数 v 是一个整数数组，表示要设置的对角线元素
func setdiagE(mat [][]int, r int, c int, v []int) {
	sort.Ints(v)
	m, n := len(mat), len(mat[0])

	i, j := r, c
	index := 0
	for ; i < m && j < n; i, j = i+1, j+1 {
		mat[i][j] = v[index]
		index++
	}
	return
}

// @lc code=end

func TestSortTheMatrixDiagonally(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[3,3,1,1],[2,2,1,2],[1,1,1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]\n
// @lcpr case=end

*/
