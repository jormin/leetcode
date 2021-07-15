package problem

import (
	"fmt"
	"strings"
)

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

// 最长回文子串
func longestPalindrome(str string) string {
	if str == "" {
		return ""
	}
	m := map[string][]int{}
	maxLength := 0
	res := string(str[0])
	for _, c := range str {
		s := string(c)
		if _, ok := m[s]; ok {
			continue
		}
		count := strings.Count(str, s)
		if count <= 1 {
			continue
		}
		fi := strings.Index(str, s)
		m[s] = []int{fi}
		for i := fi + 1; i < len(str); i++ {
			if string(str[i]) == s {
				m[s] = append(m[s], i)
				if i-m[s][0] <= maxLength {
					continue
				}
				maxLength = i - m[s][0]
				tmp := str[m[s][0] : i+1]
				if strings.Count(tmp, s) == len(tmp) {
					res = tmp
				}
				break
			}
		}
		fmt.Println(s, res)
	}
	return res
}
