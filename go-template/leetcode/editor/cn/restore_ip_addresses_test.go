/*
 * @lc app=leetcode.cn id=93 lang=golang
 * @lcpr version=30104
 *
 * [93] 复原 IP 地址
 */

package leetcode_solutions

import (
	"strconv"
	"strings"
	"testing"
)

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := []string{}
	track := []string{}

	backtrackria(s, 0, &result, &track)
	return result
}

func backtrackria(s string, start int, result *[]string, track *[]string) {
	if start == len(s) {
		if len(*track) == 4 {
			temp := strings.Join(*track, ".")
			*result = append(*result, temp)
		}

		return
	}
	for i := start; i < len(s); i++ {
		if !isValidIP(s, start, i) {
			continue
		}
		*track = append(*track, s[start:i+1])
		backtrackria(s, i+1, result, track)
		*track = (*track)[:len(*track)-1]
	}

}

func isValidIP(s string, start int, i int) bool {
	temp := s[start : i+1]
	runes := []rune(temp)
	if len(temp) >= 2 {

		if runes[0] == '0' {
			return false
		}
	}
	v, err := strconv.Atoi(temp)
	if err != nil {
		return false
	}
	if v < 0 || v > 255 {
		return false
	}
	return true

}

// @lc code=end

func TestRestoreIpAddresses(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "25525511135"\n
// @lcpr case=end

// @lcpr case=start
// "0000"\n
// @lcpr case=end

// @lcpr case=start
// "101023"\n
// @lcpr case=end

*/
