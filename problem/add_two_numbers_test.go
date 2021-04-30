package problem

import (
	"reflect"
	"testing"
)

// 测试两数相加
func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *LinkNode
		l2 *LinkNode
	}
	tests := []struct {
		name string
		args args
		want *LinkNode
	}{
		{
			name: "01",
			args: args{
				l1: makeLinkNode([]int{2, 4, 3}),
				l2: makeLinkNode([]int{5, 6, 4}),
			},
			want: makeLinkNode([]int{7, 0, 8}),
		},
		{
			name: "02",
			args: args{
				l1: makeLinkNode([]int{0}),
				l2: makeLinkNode([]int{0}),
			},
			want: makeLinkNode([]int{0}),
		}, {
			name: "03",
			args: args{
				l1: makeLinkNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: makeLinkNode([]int{9, 9, 9, 9}),
			},
			want: makeLinkNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

// 生成链表
func makeLinkNode(nums []int) *LinkNode {
	l := &LinkNode{
		Val:  0,
		Next: nil,
	}
	current := l
	for _, num := range nums {
		current.Next = &LinkNode{
			Val:  num,
			Next: nil,
		}
		current = current.Next
	}
	return l.Next
}
