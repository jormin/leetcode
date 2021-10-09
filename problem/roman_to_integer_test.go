package problem

import "testing"

// TestRomanToInt 测试罗马数字转整数
func TestRomanToInt(t *testing.T) {
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
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "02",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "03",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "04",
			args: args{s: "LVIII"},
			want: 58,
		},
		{
			name: "05",
			args: args{s: "MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := romanToInt(tt.args.s); got != tt.want {
					t.Errorf("romanToInt() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
