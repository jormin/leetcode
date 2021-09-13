package problem

import "testing"

// TestLongestPalindrome 测试最长回文子串
func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "02",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "03",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "04",
			args: args{s: "ac"},
			want: "a",
		},
		{
			name: "05",
			args: args{s: "bb"},
			want: "bb",
		},
		{
			name: "06",
			args: args{s: "bbb"},
			want: "bbb",
		},
		{
			name: "07",
			args: args{s: "bbbb"},
			want: "bbbb",
		},
		{
			name: "08",
			args: args{s: "aacabdkacaa"},
			want: "aca",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := longestPalindromic(tt.args.s); got != tt.want {
					t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
