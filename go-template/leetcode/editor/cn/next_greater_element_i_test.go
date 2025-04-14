/*
 * @lc app=leetcode.cn id=496 lang=golang
 * @lcpr version=30104
 *
 * [496] 下一个更大元素 I
 */

package leetcode_solutions

import "testing"

// @lc code=start
func nextGreaterElementTable(nums2 []int) []int {
	n := len(nums2)
	res := make([]int, n)
	s := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(s) != 0 && nums2[i] >= s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = -1
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums2[i])
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	table := nextGreaterElementTable(nums2)
	tmap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		tmap[nums2[i]] = table[i]
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = tmap[nums1[i]]
	}
	return res
}

// @lc code=end

func TestNextGreaterElementI(t *testing.T) {
	// your test code here
	nums1 := []int{2}
	nums2 := []int{2, 1}
	nextGreaterElement(nums1, nums2)

}

/*
// @lcpr case=start
// [4,1,2]\n[1,3,4,2].\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n[1,2,3,4].\n
// @lcpr case=end

*/
