package problem

import (
	"fmt"
	"reflect"
	"testing"
)

// first 第一个链表
var first = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	},
}

// firstRes 第一个链表处理后的结果
var firstRes = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	},
}

// second 第二个链表
var second = &ListNode{
	Val:  1,
	Next: nil,
}

// secondRes 第二个链表处理后的结果
var secondRes *ListNode = nil

// third 第三个链表
var third = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val:  2,
		Next: nil,
	},
}

// thirdRes 第三个链表处理后的结果
var thirdRes = &ListNode{
	Val:  1,
	Next: nil,
}

// TestGetLength 测试删除链表的倒数第N个结点
func TestGetLength(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{
			name:       "01",
			args:       args{head: first},
			wantLength: 5,
		},
		{
			name:       "02",
			args:       args{head: second},
			wantLength: 1,
		},
		{
			name:       "03",
			args:       args{head: third},
			wantLength: 2,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotLength := getLength(tt.args.head); gotLength != tt.wantLength {
					t.Errorf("getLength() = %v, want %v", gotLength, tt.wantLength)
				}
			},
		)
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{head: first, n: 2},
			want: firstRes,
		},
		{
			name: "02",
			args: args{head: second, n: 1},
			want: secondRes,
		},
		{
			name: "03",
			args: args{head: third, n: 1},
			want: thirdRes,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := removeNthFromEnd(tt.args.head, tt.args.n)
				fmt.Println(tt.args.head, tt.args.n, tt.want, got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
