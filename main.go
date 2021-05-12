package main

import "fmt"

func main() {
	fmt.Printf("%v,%v,%v,%v,%v,%v\n", []byte(`a`), []byte(`z`), []byte(`0`), []byte(`9`), []byte(` `), []byte(`-`))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 字符对应的数组，默认所有位置都为false，256个已足够包含题目中的字符
	bitSet := [256]bool{}
	left, right, length := 0, 0, 0
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
		cur := right - left
		if length < cur {
			length = cur
		}
		// 如果左侧位置加最新的最大长度大于等于字符总长度，说明后续不可能再有比这个长度更大的连续无重复子串了
		// 如果右侧位置已经到头则需要停止循环
		if left+length >= len(s) || right >= len(s) {
			break
		}
	}
	return length
}
