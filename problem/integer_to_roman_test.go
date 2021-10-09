package problem

import "testing"

// TestIntToRoman 测试整数转罗马数字
func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "02",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "03",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "04",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "05",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := intToRoman(tt.args.num); got != tt.want {
					t.Errorf("intToRoman() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
