/*
 * @lc app=leetcode.cn id=75 lang=golang
 * @lcpr version=30100
 *
 * [75] 颜色分类
 */

package leetcode_solutions

import "testing"

// @lc code=start
func sortColors2(nums []int) {
	if len(nums) == 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return

}
func sortColors3(nums []int) {
	p0, p2 := 0, len(nums)-1
	p := 0
	for p <= p2 {
		if nums[p] == 0 {
			swap(nums, p0, p)
			p0++
		} else if nums[p] == 2 {
			swap(nums, p, p2)
			p2--
		} else if nums[p] == 1 {
			p++
		}
		if p < p0 {
			p = p0
		}
	}
	return

}
func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
	return
}

func sortColors(nums []int) {
	//把0放在前面
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			swap(nums, slow, fast)
			slow++
		}
		fast++
	}
	fast = slow

	for fast < len(nums) {
		if nums[fast] == 1 {
			swap(nums, slow, fast)
			slow++
		}
		fast++
	}
	return

}

// @lc code=end

func TestSortColors(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

*/
