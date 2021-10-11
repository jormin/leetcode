package problem

import "testing"

// TestLongestCommonPrefix 测试最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "02",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := longestCommonPrefix(tt.args.strs); got != tt.want {
					t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
