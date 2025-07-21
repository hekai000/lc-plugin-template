/*
 * @lc app=leetcode.cn id=127 lang=golang
 * @lcpr version=30104
 *
 * [127] 单词接龙
 */

package leetcode_solutions

import "testing"

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	validWordMap := make(map[string]bool)
	for _, w := range wordList {
		validWordMap[w] = true
	}
	if !validWordMap[endWord] {
		return 0
	}
	queue := []string{beginWord}
	visitedWords := make(map[string]bool)
	visitedWords[beginWord] = true
	step := 1
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			curWord := queue[0]
			queue = queue[1:]
			if curWord == endWord {
				return step
			}
			for _, nextWord := range transformWord(curWord, validWordMap) {
				if !visitedWords[nextWord] && validWordMap[nextWord] {
					queue = append(queue, nextWord)
					visitedWords[nextWord] = true
				}
			}
		}
		step++
	}
	return 0
}

func transformWord(word string, validWordMap map[string]bool) []string {
	letters := "abcdefghijklmnopqrstuvwxyz" // 可变换的字母表
	results := []string{}
	runes := []rune(word) // 转为 rune 切片以支持 Unicode

	for i := 0; i < len(runes); i++ {
		original := runes[i] // 备份原始字符
		for _, char := range letters {
			if char == original { // 跳过原始字符
				continue
			}
			runes[i] = char          // 修改第 i 位字符
			newWord := string(runes) // 生成新字符串
			if validWordMap[newWord] {
				results = append(results, newWord)
			}

		}
		runes[i] = original // 恢复原字符
	}
	return results
}

// @lc code=end

func TestWordLadder(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]\n
// @lcpr case=end

// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log"]\n
// @lcpr case=end

*/
