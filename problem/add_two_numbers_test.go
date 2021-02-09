package problem

import (
	"reflect"
	"testing"
)

// 测试两数相加
func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{
				l1: makeListNode([]int{2, 4, 3}),
				l2: makeListNode([]int{5, 6, 4}),
			},
			want: makeListNode([]int{7, 0, 8}),
		},
		{
			name: "02",
			args: args{
				l1: makeListNode([]int{0}),
				l2: makeListNode([]int{0}),
			},
			want: makeListNode([]int{0}),
		}, {
			name: "03",
			args: args{
				l1: makeListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: makeListNode([]int{9, 9, 9, 9}),
			},
			want: makeListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("addTwoNumbers() = %v, want %v", got.Info(), tt.want.Info())
				}
			},
		)
	}
}

// 生成链表
func makeListNode(nums []int) *ListNode {
	l := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := l
	for _, num := range nums {
		current.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		current = current.Next
	}
	return l.Next
}
