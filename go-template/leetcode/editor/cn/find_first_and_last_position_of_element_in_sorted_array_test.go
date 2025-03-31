/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=30104
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

package leetcode_solutions

import "testing"

// @lc code=start
func searchRange(nums []int, target int) []int {
	leftB := leftBound(nums, target)
	rightB := rightBound(nums, target)
	return []int{leftB, rightB}
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left < 0 || left >= len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right < 0 || right >= len(nums) {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}

// @lc code=end

func TestFindFirstAndLastPositionOfElementInSortedArray(t *testing.T) {
	// your test code here
	a := []int{1}
	searchRange(a, 0)
}

/*
// @lcpr case=start
// [5,7,7,8,8,10]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
