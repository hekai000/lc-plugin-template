/*
 * @lc app=leetcode.cn id=901 lang=golang
 * @lcpr version=30104
 *
 * [901] 股票价格跨度
 */

package leetcode_solutions

import (
	"fmt"
	"testing"
)

// @lc code=start
type StockSpanner struct {
	stk [][2]int
}

func Constructorpp() StockSpanner {
	return StockSpanner{stk: make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	count := 1
	for len(this.stk) > 0 && price >= this.stk[len(this.stk)-1][0] {
		prev := this.stk[len(this.stk)-1]
		this.stk = this.stk[:len(this.stk)-1]
		count += prev[1]
	}
	this.stk = append(this.stk, [2]int{price, count})
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

func TestOnlineStockSpan(t *testing.T) {
	// your test code here
	obj := Constructorpp()
	param_1 := obj.Next(100)
	fmt.Println(param_1)
	param_2 := obj.Next(80)
	fmt.Println(param_2)
	param_3 := obj.Next(60)
	fmt.Println(param_3)
	param_4 := obj.Next(70)
	fmt.Println(param_4)
	param_5 := obj.Next(60)
	fmt.Println(param_5)
	param_6 := obj.Next(75)
	fmt.Println(param_6)
	param_7 := obj.Next(85)
	fmt.Println(param_7)

}

/*
// @lcpr case=start
// ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]\n[[], [100], [80], [60], [70], [60], [75], [85]]\n
// @lcpr case=end

*/
