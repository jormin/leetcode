package problem

import (
	"fmt"
	"strconv"
	"strings"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 
// 示例 1：
// 输入：x = 123
// 输出：321
//
// 示例 2：
// 输入：x = -123
// 输出：-321
//
// 示例 3：
// 输入：x = 120
// 输出：21
//
// 示例 4：
// 输入：x = 0
// 输出：0
// 
// 提示：
// -231 <= x <= 231 - 1
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-integer
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// reverse 整数反转
func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x *= -1
	}
	s := fmt.Sprintf("%d", x)
	l := len(s)
	arr := make([]string, l)
	for i, c := range s {
		arr[l-1-i] = fmt.Sprintf("%c", c)
	}
	res, _ := strconv.Atoi(strings.Join(arr, ""))
	res *= flag
	if res < -1*1<<31 || res > 1<<31+1 {
		res = 0
	}
	return res
}
