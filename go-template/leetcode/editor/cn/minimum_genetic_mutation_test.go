/*
 * @lc app=leetcode.cn id=433 lang=golang
 * @lcpr version=30104
 *
 * [433] 最小基因变化
 */

package leetcode_solutions

import "testing"

// @lc code=start
func minMutation(startGene string, endGene string, bank []string) int {
	queue := []string{startGene}
	visited := make(map[string]bool)
	step := 0
	visited[startGene] = true
	validGene := make(map[string]bool)
	for _, v := range bank {
		validGene[v] = true
	}
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == endGene {
				return step
			}
			for _, neigbhor := range getNeighborsmm(cur) {
				if !visited[neigbhor] && validGene[neigbhor] {
					queue = append(queue, neigbhor)
					visited[neigbhor] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNeighborsmm(cur string) []string {
	neighbors := make([]string, 0)
	runes := []rune(cur)
	nmap1 := make(map[rune][]rune)
	nmap1['A'] = []rune{'G', 'T', 'C'}
	nmap1['G'] = []rune{'A', 'T', 'C'}
	nmap1['T'] = []rune{'G', 'A', 'C'}
	nmap1['C'] = []rune{'G', 'T', 'A'}
	for i := 0; i < 8; i++ {
		ch := runes[i]
		for _, v := range nmap1[ch] {
			runes[i] = v
			neighbors = append(neighbors, string(runes))
			runes[i] = ch
		}
	}
	return neighbors
}

// @lc code=end

func TestMinimumGeneticMutation(t *testing.T) {
	// your test code here
	startG := "AACCGGTT"
	endG := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	minMutation(startG, endG, bank)

}

/*
// @lcpr case=start
// "AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AACCGGTT"\n"AAACGGTA"\n["AACCGGTA","AACCGCTA","AAACGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AAAAACCC"\n"AACCCCCC"\n["AAAACCCC","AAACCCCC","AACCCCCC"]\n
// @lcpr case=end

*/
