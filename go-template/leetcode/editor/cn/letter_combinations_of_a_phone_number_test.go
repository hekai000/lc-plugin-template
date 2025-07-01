/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=30104
 *
 * [17] 电话号码的字母组合
 */

package leetcode_solutions

import "testing"

// @lc code=start
func letterCombinations(digits string) []string {
	result := []string{}
	track := []rune{}
	if len(digits) == 0 {
		return []string{}
	}
	phoneMap := make(map[string]string, 8)
	phoneMap["2"] = "abc"
	phoneMap["3"] = "def"
	phoneMap["4"] = "ghi"
	phoneMap["5"] = "jkl"
	phoneMap["6"] = "mno"
	phoneMap["7"] = "pqrs"
	phoneMap["8"] = "tuv"
	phoneMap["9"] = "wxyz"

	backtracklc(digits, 0, &result, &track, phoneMap)

	return result
}

func backtracklc(digits string, start int, result *[]string, track *[]rune, phoneMap map[string]string) {
	if start == len(digits) {
		*result = append(*result, string(*track))
		return
	}
	choicList := phoneMap[string(([]rune(digits))[start])]
	for _, ch := range []rune(choicList) {
		*track = append(*track, ch)
		backtracklc(digits, start+1, result, track, phoneMap)
		*track = (*track)[:len(*track)-1]
	}
}

// @lc code=end

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
