/*
 * @lc app=leetcode.cn id=71 lang=golang
 * @lcpr version=30104
 *
 * [71] 简化路径
 */

package leetcode_solutions

import (
	"strings"
	"testing"
)

// @lc code=start

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stk := []string{}
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
			continue
		}
		stk = append(stk, part)
	}
	res := ""
	for _, dir := range stk {
		res += "/" + dir
	}
	if res == "" {
		return "/"
	}
	return res
}

// type MyStack2 struct {
// 	q        []string
// 	top_elem string
// }

// func Constructor3() MyStack2 {
// 	return MyStack2{}
// }

// func (this *MyStack2) Push(x string) {
// 	this.q = append(this.q, x)
// 	this.top_elem = x
// }

// func (this *MyStack2) Pop() string {
// 	size := len(this.q)
// 	for size > 2 {
// 		this.q = append(this.q, this.q[0])
// 		this.q = this.q[1:]
// 		size--
// 	}
// 	this.top_elem = this.q[0]
// 	this.q = append(this.q, this.q[0])
// 	this.q = this.q[1:]

// 	result := this.q[0]
// 	this.q = this.q[1:]
// 	return result
// }

// func (this *MyStack2) Top() string {
// 	return this.top_elem
// }

// func (this *MyStack2) Lens() int {
// 	return len(this.q)
// }

// func (this *MyStack2) Print() string {
// 	res := ""
// 	for _, v := range this.q {
// 		res += v
// 	}
// 	return res
// }

// func (this *MyStack2) Empty() bool {
// 	return len(this.q) == 0
// }

// func simplifyPath2(path string) string {
// 	stack := Constructor3()
// 	arr := splitPath(path)
// 	for _, v := range arr {
// 		if !stack.Empty() {
// 			top := stack.Top()

// 			if top == "/" && string(v) == "/" {
// 				continue
// 			} else if string(v) == ".." {
// 				stack.Pop()
// 				if stack.Empty() {
// 					continue
// 				}
// 				stack.Pop()
// 				continue
// 			} else if string(v) == "." {
// 				continue
// 			}
// 		}

// 		stack.Push(v)

// 	}
// 	if stack.Top() == "/" {
// 		if stack.Lens() > 1 {
// 			stack.Pop()
// 		}

// 	}
// 	if stack.Empty() {
// 		stack.Push("/")
// 	}
// 	res := stack.Print()
// 	return res

// }
// func splitPath(path string) []string {
// 	var result []string // 存储最终结果
// 	var buffer []rune   // 临时积累非分隔符字符
// 	separator := '/'    // 分隔符

// 	for _, char := range path {
// 		if char == separator {
// 			// 遇到分隔符：将积累的非分隔符加入结果，并添加分隔符
// 			if len(buffer) > 0 {
// 				result = append(result, string(buffer))
// 				buffer = buffer[:0] // 清空缓冲区
// 			}
// 			result = append(result, string(separator))
// 		} else {
// 			// 非分隔符：积累到缓冲区
// 			buffer = append(buffer, char)
// 		}
// 	}
// 	// 处理末尾剩余的非分隔符
// 	if len(buffer) > 0 {
// 		result = append(result, string(buffer))
// 	}
// 	return result
// }

// @lc code=end

func TestSimplifyPath(t *testing.T) {
	// your test code here
	a := "/.."
	simplifyPath(a)

}

/*
// @lcpr case=start
// "/home/"\n
// @lcpr case=end

// @lcpr case=start
// "/home//foo/"\n
// @lcpr case=end

// @lcpr case=start
// "/home/user/Documents/../Pictures"\n
// @lcpr case=end

// @lcpr case=start
// "/../"\n
// @lcpr case=end

// @lcpr case=start
// "/.../a/../b/c/../d/./"\n
// @lcpr case=end

*/
