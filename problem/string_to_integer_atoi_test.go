package problem

import (
	"testing"
)

// TestMyAtoi 测试字符串转换整数
func TestMyAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "02",
			args: args{s: "+42"},
			want: 42,
		},
		{
			name: "03",
			args: args{s: "-42"},
			want: -42,
		},
		{
			name: "04",
			args: args{s: " -42"},
			want: -42,
		},
		{
			name: "05",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "06",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "07",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "08",
			args: args{s: "2147483648"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := myAtoi(tt.args.s); got != tt.want {
					t.Errorf("myAtoi() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
