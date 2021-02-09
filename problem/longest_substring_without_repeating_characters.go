package problem

// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
// 
// 示例1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//
// 示例 4:
// 输入: s = ""
// 输出: 0
// 
// 提示：
// 
// 0 <= s.length <= 5 * 104
// s由英文字母、数字、符号和空格组成
// 
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 无重复字符的最长子串 - 位图
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	bitSet := [256]bool{}
	left, right, res := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if res < right-left {
			res = right - left
		}
		if left+res >= len(s) || right >= len(s) {
			break
		}
	}
	return res
}
