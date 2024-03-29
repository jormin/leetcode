package problem

import "testing"

// TestReverse 测试整数反转
func TestReverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{x: 123},
			want: 321,
		},
		{
			name: "02",
			args: args{x: -123},
			want: -321,
		},
		{
			name: "03",
			args: args{x: 120},
			want: 21,
		},
		{
			name: "04",
			args: args{x: 0},
			want: 0,
		},
		{
			name: "05",
			args: args{x: -1 * 1 << 31},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := reverse(tt.args.x); got != tt.want {
					t.Errorf("reverse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
