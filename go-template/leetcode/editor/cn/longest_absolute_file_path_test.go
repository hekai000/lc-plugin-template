/*
 * @lc app=leetcode.cn id=388 lang=golang
 * @lcpr version=30104
 *
 * [388] 文件的最长绝对路径
 */

package leetcode_solutions

import "testing"

// @lc code=start
func lengthLongestPath(input string) int {
	st := []int{}
	ans := 0
	i, n := 0, len(input)
	for i < n {
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}
		length, isFile := 0, false

		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}

		i++

		for len(st) >= depth {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			length += st[len(st)-1] + 1
		}
		if isFile {
			ans = max(ans, length)
		} else {
			st = append(st, length)
		}

	}
	return ans
}

// @lc code=end

func TestLongestAbsoluteFilePath(t *testing.T) {
	// your test code here
	a := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	lengthLongestPath(a)
}

/*
// @lcpr case=start
// "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"\n
// @lcpr case=end

// @lcpr case=start
// "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

// @lcpr case=start
// "file1.txt\nfile2.txt\nlongfile.txt"\n
// @lcpr case=end

*/
