package problem

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "02",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "03",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "04",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
					t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
