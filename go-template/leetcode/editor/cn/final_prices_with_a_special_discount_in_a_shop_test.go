/*
 * @lc app=leetcode.cn id=1475 lang=golang
 * @lcpr version=30104
 *
 * [1475] 商品折扣后的最终价格
 */

package leetcode_solutions

import "testing"

// @lc code=start
func finalPrices(prices []int) []int {
	n := len(prices)

	res := make([]int, n)
	stk := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] > prices[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = 0
		} else {
			res[i] = stk[len(stk)-1]
		}
		stk = append(stk, prices[i])
	}

	for i := 0; i < n; i++ {
		prices[i] = prices[i] - res[i]
	}
	return prices

}

// @lc code=end

func TestFinalPricesWithASpecialDiscountInAShop(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [8,4,6,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [10,1,1,6]\n
// @lcpr case=end

*/
