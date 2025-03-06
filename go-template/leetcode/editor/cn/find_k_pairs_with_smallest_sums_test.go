/*
 * @lc app=leetcode.cn id=373 lang=golang
 * @lcpr version=30100
 *
 * [373] 查找和最小的 K 对数字
 */

package leetcode_solutions

import (
	"container/heap"
	"testing"
)

// @lc code=start
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return (pq[i][0] + pq[i][1]) < (pq[j][0] + pq[j][1])
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int{nums1[i], nums2[0], 0})
	}

	res := [][]int{}

	for pq.Len() > 0 && k > 0 {
		cur := heap.Pop(pq).([]int)
		k--

		nextIndex := cur[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})

		}
		res = append(res, []int{cur[0], cur[1]})
	}
	return res

}

// @lc code=end

func TestFindKPairsWithSmallestSums(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,7,11]\n[2,4,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2]\n[1,2,3]\n2\n
// @lcpr case=end

*/
