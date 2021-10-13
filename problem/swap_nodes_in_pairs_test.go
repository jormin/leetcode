package problem

import (
	"reflect"
	"testing"
)

// TestSwapPairs 测试两两交换链表中的节点
func TestSwapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "02",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "03",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("swapPairs() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
