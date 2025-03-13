/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=30100
 *
 * [151] 反转字符串中的单词
 */

package leetcode_solutions

import (
	"fmt"
	"strings"
	"testing"
)

// @lc code=start
func reverseWords(s string) string {
	ss := reverseWord(s)

	parts := strings.Fields(ss)
	words := []string{}
	for _, word := range parts {

		words = append(words, reverseWord(word))

	}
	return strings.Join(words, " ")

}

func reverseWord(word string) string {
	i, j := 0, len(word)-1
	bs := []rune(word)
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
	return string(bs)
}

// @lc code=end

func TestReverseWordsInAString(t *testing.T) {
	// your test code here
	s := "a good  example"
	fmt.Println(reverseWords(s))

}

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// " hello world "\n
// @lcpr case=end

// @lcpr case=start
// "a good  example"\n
// @lcpr case=end

*/
