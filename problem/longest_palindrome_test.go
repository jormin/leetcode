package problem

import "testing"

// TestLongestPalindromeNum 测试最长回文串
func TestLongestPalindromeNum(t *testing.T) {
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
			args: args{s: ""},
			want: 0,
		},
		{
			name: "02",
			args: args{s: "abccccdd"},
			want: 7,
		},
		{
			name: "03",
			args: args{s: "AAAAAA"},
			want: 6,
		},
		{
			name: "04",
			args: args{s: "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := longestPalindromeNum(tt.args.s); got != tt.want {
					t.Errorf("longestPalindromeNum() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
