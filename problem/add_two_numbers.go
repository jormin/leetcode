package problem

// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
//
// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
//
// 示例 3：
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
// 提示：
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// LinkNode 链表结点
type LinkNode struct {
	Val  int
	Next *LinkNode
}

// addTwoNumbers 两数相加
func addTwoNumbers(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	res := &LinkNode{
		Val:  0,
		Next: nil,
	}
	n1, n2, carry, current := 0, 0, 0, res
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &LinkNode{
			Val:  (n1 + n2 + carry) % 10,
			Next: nil,
		}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return res.Next
}
