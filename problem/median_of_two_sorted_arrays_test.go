package problem

import "testing"

// TestFindMedianSortedArrays 测试寻找两个正序数组的中位数
func TestFindMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "01",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.00000,
		},
		{
			name: "02",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.50000,
		},
		{
			name: "03",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0.00000,
		},
		{
			name: "04",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.00000,
		},
		{
			name: "05",
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.00000,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
					t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
