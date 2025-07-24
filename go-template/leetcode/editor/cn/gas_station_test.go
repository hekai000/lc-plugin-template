/*
 * @lc app=leetcode.cn id=134 lang=golang
 * @lcpr version=30104
 *
 * [134] 加油站
 */

package leetcode_solutions

import "testing"

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	curGas := 0
	startStation := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]

		curGas += gas[i] - cost[i]
		if curGas < 0 {
			startStation = i + 1
			curGas = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return startStation
}
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		if checkComplete(gas, cost, i) {
			return i
		}
	}
	return -1
}

func checkComplete(gas []int, cost []int, i int) bool {
	n := len(gas)
	gas = append(gas, gas...)
	cost = append(cost, cost...)
	totalGas := 0
	for j := i; j < i+n; j++ {
		totalGas += gas[j]
		if totalGas < cost[j] {
			return false
		}
		totalGas -= cost[j]
	}
	return true
}

// @lc code=end

func TestGasStation(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n[3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[3,4,3]\n
// @lcpr case=end

*/
