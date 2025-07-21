/*
 * @lc app=leetcode.cn id=354 lang=golang
 * @lcpr version=30104
 *
 * [354] 俄罗斯套娃信封问题
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1] //第二个元素降序
		}
		return envelopes[i][0] < envelopes[j][0] //第一个元素升序
	})
	height := make([]int, n)
	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}
	return lengthOfLISRUS(height)
}

func lengthOfLISRUS2(height []int) int {
	dpa := make([]int, len(height))
	for i := range dpa {
		dpa[i] = 1
	}
	for i := 0; i < len(height); i++ {
		for j := 0; j < i; j++ {
			if height[j] < height[i] {
				dpa[i] = maxrus(dpa[i], 1+dpa[j])
			}
		}
	}
	res := 0
	for _, v := range dpa {
		res = maxrus(res, v)
	}
	return res
}
func lengthOfLISRUS(height []int) int {
	n := len(height)
	top := make([]int, n)
	piles := 0
	for _, poker := range height {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] >= poker {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func maxrus(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestRussianDollEnvelopes(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[5,4],[6,4],[6,7],[2,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[1,1],[1,1]]\n
// @lcpr case=end

*/
