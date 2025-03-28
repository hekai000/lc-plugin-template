/*
 * @lc app=leetcode.cn id=220 lang=golang
 * @lcpr version=30104
 *
 * [220] 存在重复元素 III
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	} else {
		return (x+1)/w - 1
	}
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {

	window := make(map[int]int)
	w := valueDiff + 1

	for i := 0; i < len(nums); i++ {
		m := getID(nums[i], w)
		if _, ok := window[m]; ok {
			return true
		}
		if v, ok := window[m-1]; ok && myAbs(nums[i]-v) < w {
			return true
		}
		if v, ok := window[m+1]; ok && myAbs(v-nums[i]) < w {
			return true
		}

		window[m] = nums[i]

		if i >= indexDiff {
			delete(window, getID(nums[i-indexDiff], w))
		}
	}
	return false
}
func containsNearbyAlmostDuplicate2(nums []int, indexDiff int, valueDiff int) bool {
	left, right := 0, 0
	window := []int{}

	for right < len(nums) {
		c := nums[right]
		right++

		if existsAbsLessThanK(window, c, valueDiff) {
			return true
		}
		window = append(window, c)

		for right-left > indexDiff {
			// d := nums[left]
			left++

			window = window[1:]
		}

	}
	return false
}

func existsAbsLessThanK(arr []int, target int, k int) bool {
	arrSlice := make([]int, len(arr))
	copy(arrSlice, arr[:])
	sort.Ints(arrSlice)
	lower := target - k
	upper := target + k
	// 查找第一个大于等于lower的位置
	left := sort.SearchInts(arrSlice, lower)
	// 若存在元素且在upper范围内则返回true
	if left < len(arr) && arrSlice[left] <= upper {
		return true
	}
	return false
}

func checkD(window []int, c int, valuvalueDiff int) bool {
	for _, v := range window {
		if myAbs(v-c) <= valuvalueDiff {
			return true
		}
	}
	return false
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

func TestContainsDuplicateIii(t *testing.T) {
	// your test code here
	a := []int{1, 5, 9, 1, 5, 9}

	containsNearbyAlmostDuplicate(a, 2, 3)

}

/*
// @lcpr case=start
// [1,2,3,1]\n3\n0\n
// @lcpr case=end

// @lcpr case=start
// [1,5,9,1,5,9]\n2\n3\n
// @lcpr case=end

*/
