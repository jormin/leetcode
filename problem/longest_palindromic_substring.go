package problem

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 
// 示例 1：
// 
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 
// 输入：s = "cbbd"
// 输出："bb"
// 示例 3：
// 
// 输入：s = "a"
// 输出："a"
// 示例 4：
// 
// 输入：s = "ac"
// 输出："a"
//
// 提示：
// 
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
// 
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// longestPalindromic 最长回文子串
func longestPalindromic(str string) string {
	if str == "" {
		return ""
	}
	var res string
	m := 0
	l := len(str)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			s := str[i:j]
			if isPalindromic(s) && len(s) > m {
				m = len(s)
				res = s
			}
		}
	}
	return res
}

// isPalindromic 是否是回文串
func isPalindromic(str string) bool {
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}
