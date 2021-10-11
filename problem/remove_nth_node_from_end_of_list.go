package problem

import (
	"fmt"
	"strings"
)

// 给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
//
// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
//
// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]
//
// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]
//
// 提示：
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// String 字符串
func (l *ListNode) String() string {
	var arr []string
	for ; l != nil; l = l.Next {
		arr = append(arr, fmt.Sprintf("%d", l.Val))
	}
	return strings.Join(arr, ",")
}

// removeNthFromEnd 删除链表的倒数第N个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getLength(head)
	var node = &ListNode{
		Val:  0,
		Next: head,
	}
	cur := node
	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return node.Next
}

// getLength 获取链表的长度
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}
