package problem

// 数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 有效括号组合需满足：左括号必须以正确的顺序闭合。
//
// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
// 示例 2：
// 输入：n = 1
// 输出：["()"]
//
// 提示：
// 1 <= n <= 8
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/generate-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// generateParenthesis 括号生成
func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	var res []string
	res = append(res)
	dfs("", n, 0, 0, &res)
	return res
}

// 深度优先搜索算法
func dfs(path string, n int, left int, right int, res *[]string) {
	if left == n && right == n {
		*res = append(*res, path)
	}
	if left < n {
		dfs(path+"(", n, left+1, right, res)
	}
	if right < left {
		dfs(path+")", n, left, right+1, res)
	}
}
