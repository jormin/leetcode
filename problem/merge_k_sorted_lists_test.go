package problem

import (
	"reflect"
	"testing"
)

// TestMergeKLists 测试合并K个升序链表
func TestMergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "01",
			args: args{
				lists: []*ListNode{
					&ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
					&ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
					&ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "02",
			args: args{lists: []*ListNode{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
