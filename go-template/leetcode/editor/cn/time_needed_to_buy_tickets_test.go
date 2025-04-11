/*
 * @lc app=leetcode.cn id=2073 lang=golang
 * @lcpr version=30104
 *
 * [2073] 买票需要的时间
 */

package leetcode_solutions

import "testing"

// @lc code=start
func timeRequiredToBuy(tickets []int, k int) int {
	count := 0
	for tickets[k] != 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] == 0 {
				continue
			} else {
				tickets[i]--
				count++
				if tickets[k] == 0 {
					break
				}
			}

		}
	}
	return count
}

// @lc code=end

func TestTimeNeededToBuyTickets(t *testing.T) {
	// your test code here
	tik := []int{84, 49, 5, 24, 70, 77, 87, 8}
	timeRequiredToBuy(tik, 3)
}

/*
// @lcpr case=start
// [2,3,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [5,1,1,1]\n0\n
// @lcpr case=end

*/
