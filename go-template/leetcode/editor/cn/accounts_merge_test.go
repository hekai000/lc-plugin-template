/*
 * @lc app=leetcode.cn id=721 lang=golang
 * @lcpr version=30104
 *
 * [721] 账户合并
 */

package leetcode_solutions

import (
	"sort"
	"testing"
)

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	emailToindexs := make(map[string][]int)
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			emailToindexs[email] = append(emailToindexs[email], i)
		}
	}

	var res [][]string
	visitedEmails := make(map[string]bool)

	for email := range emailToindexs {
		if visitedEmails[email] {
			continue
		}
		var mergeEmails []string
		queue := []string{email}
		visitedEmails[email] = true
		for len(queue) > 0 {
			curEmail := queue[0]
			queue = queue[1:]
			mergeEmails = append(mergeEmails, curEmail)
			for _, index := range emailToindexs[curEmail] {
				account := accounts[index]
				for j := 1; j < len(account); j++ {
					nextMail := account[j]
					if !visitedEmails[nextMail] {
						queue = append(queue, nextMail)
						visitedEmails[nextMail] = true
					}
				}
			}
		}
		userName := accounts[emailToindexs[email][0]][0]
		sort.Strings(mergeEmails)
		mergeEmails = append([]string{userName}, mergeEmails...)
		res = append(res, mergeEmails)
	}
	return res

}

// @lc code=end

func TestAccountsMerge(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John",\n"johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]\n
// @lcpr case=end

// @lcpr case=start
// \n[["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]\n
// @lcpr case=end

*/
