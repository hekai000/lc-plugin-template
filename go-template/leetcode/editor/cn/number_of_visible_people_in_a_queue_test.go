/*
 * @lc app=leetcode.cn id=1944 lang=golang
 * @lcpr version=30104
 *
 * [1944] 队列中可以看到的人数
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stk := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		count := 0
		for len(stk) > 0 && heights[i] > stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
			count++
		}
		if len(stk) == 0 {
			res[i] = count
		} else {
			res[i] = count + 1
		}
		stk = append(stk, heights[i])
	}
	return res
}

// @lc code=end

func TestNumberOfVisiblePeopleInAQueue(t *testing.T) {
	// your test code here
	heights := []int{10, 6, 8, 5, 11, 9}
	canSeePersonsCount(heights)

}

/*
// @lcpr case=start
// [10,6,8,5,11,9]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,2,3,10]\n
// @lcpr case=end

*/
