/*
 * @lc app=leetcode.cn id=378 lang=golang
 * @lcpr version=30100
 *
 * [378] 有序矩阵中第 K 小的元素
 */

package leetcode_solutions

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

// @lc code=start
func kthSmallest2(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	arr := make([]int, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			arr = append(arr, matrix[i][j])
		}
	}
	sort.Ints(arr)
	return arr[k-1]
}

type PriorityQueuePQ [][]int

func (pq PriorityQueuePQ) Len() int {
	return len(pq)
}

func (pq PriorityQueuePQ) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueuePQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueuePQ) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueuePQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
func kthSmallest(matrix [][]int, k int) int {

	pq := &PriorityQueuePQ{}
	heap.Init(pq)

	for i := range matrix {
		heap.Push(pq, []int{matrix[i][0], i, 0})
	}
	res := -1

	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		res = cur[0]
		k--
		i, j := cur[1], cur[2]
		if j+1 < len(matrix[i]) {
			heap.Push(pq, []int{matrix[i][j+1], i, j + 1})

		}
	}
	return res

}

// @lc code=end

func TestKthSmallestElementInASortedMatrix(t *testing.T) {
	// your test code here
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	a := kthSmallest(matrix, 8)
	fmt.Println(a)
}

/*
// @lcpr case=start
// [[1,5,9],[10,11,13],[12,13,15]]\n8\n
// @lcpr case=end

// @lcpr case=start
// [[-5]]\n1\n
// @lcpr case=end

*/
