package problem

import (
	"reflect"
	"testing"
)

// TestThreeSum 三数之和
func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "01",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "02",
			args: args{nums: []int{}},
			want: nil,
		},
		{
			name: "03",
			args: args{nums: []int{0}},
			want: nil,
		},
		{
			name: "04",
			args: args{nums: []int{-1, 0, 1, 1, 2, 2, -1, -4}},
			want: [][]int{
				{-4, 2, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("threeSum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
