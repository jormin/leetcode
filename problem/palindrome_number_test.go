package problem

import "testing"

// TestIsPalindrome 测试回文数
func TestIsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{x: 121},
			want: true,
		},
		{
			name: "02",
			args: args{x: -121},
			want: false,
		},
		{
			name: "03",
			args: args{x: 0},
			want: true,
		},
		{
			name: "04",
			args: args{x: 10},
			want: false,
		},
		{
			name: "05",
			args: args{x: -101},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isPalindrome(tt.args.x); got != tt.want {
					t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
