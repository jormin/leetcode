package problem

// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
// 示例 1：
// 输入：s = "()"
// 输出：true
//
// 示例2：
// 输入：s = "()[]{}"
// 输出：true
//
// 示例3：
// 输入：s = "(]"
// 输出：false
//
// 示例4：
// 输入：s = "([)]"
// 输出：false
//
// 示例5：
// 输入：s = "{[]}"
// 输出：true
//
// 提示：
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// isValid 有效的括号
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	// map存储右括号为key，左括号为值
	m := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; ok {
			// 取到的字符是map存在的元素，说明取到的是右括号
			// 如果取到了右括号，但是栈为空，或者栈取的最后入的符号不是对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			// 匹配到的字符出栈
			stack = stack[:len(stack)-1]
		} else {
			// 取到的是map不存在的元素，说明是左括号，入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
