package problem

import "testing"

// TestThreeSumClosest 测试最接近的三数之和
func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "02",
			args: args{
				nums:   []int{-100, -98, -2, -1},
				target: -101,
			},
			want: -101,
		},
		{
			name: "03",
			args: args{
				nums:   []int{-1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "04",
			args: args{
				nums:   []int{-100, -100, -98, -2, -1},
				target: -101,
			},
			want: -101,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
