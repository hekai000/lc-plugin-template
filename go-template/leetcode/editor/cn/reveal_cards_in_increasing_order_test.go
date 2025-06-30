/*
 * @lc app=leetcode.cn id=950 lang=golang
 * @lcpr version=30104
 *
 * [950] 按递增顺序显示卡牌
 */

package leetcode_solutions

import (
	"container/list"
	"sort"
	"testing"
)

// @lc code=start
func deckRevealedIncreasing(deck []int) []int {
	res := list.New()
	sort.Ints(deck)
	n := len(deck)

	for i := n - 1; i >= 0; i-- {

		if res.Len() > 0 {
			e := res.Back()
			res.Remove(e)
			res.PushFront(e.Value.(int))
		}
		res.PushFront(deck[i])

	}
	arr := make([]int, n)
	index := 0
	for e := res.Front(); e != nil; e = e.Next() {
		arr[index] = e.Value.(int)
		index++
	}
	return arr
}

// @lc code=end

func TestRevealCardsInIncreasingOrder(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [17,13,11,2,3,5,7]\n
// @lcpr case=end

*/
