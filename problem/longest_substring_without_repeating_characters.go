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
	// 位图，默认所有位置都为false，256个已足够包含题目中的字符
	bitSet := [256]bool{}
	left, right, res := 0, 0, 0
	for left < len(s) {
		// 判断右侧位置字符对应的值是否为true
		if bitSet[s[right]] {
			// 如果值为true，证明这个字符已经存在，则需要左侧位进1
			bitSet[s[left]] = false
			left++
		} else {
			// 如果值为false，说明这个字符当前尚不存在，需要设置为true，并将右侧位进1
			bitSet[s[right]] = true
			right++
		}
		// 判断完后需要计算长度
		if res < right-left {
			res = right - left
		}
		// 如果左侧位置加最新的最大长度大于等于字符总长度，说明后续不可能再有比这个长度更大的连续无重复子串了
		// 如果右侧位置已经到头则需要停止循环
		if left+res >= len(s) || right >= len(s) {
			break
		}
	}
	return res
}
