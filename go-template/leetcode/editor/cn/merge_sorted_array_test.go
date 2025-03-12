/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=30100
 *
 * [88] 合并两个有序数组
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//两个指针在数组末尾
	i, j, p := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}
	for j >= 0 {
		nums1[p] = nums2[j]
		j--
		p--
	}
	return

}

// @lc code=end

func TestMergeSortedArray(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/
