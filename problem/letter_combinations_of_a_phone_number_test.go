package problem

import (
	"reflect"
	"testing"
)

// TestLetterCombinations 测试电话号码的字母组合
func TestLetterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "01",
			args: args{digits: "23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "02",
			args: args{digits: ""},
			want: nil,
		},
		{
			name: "03",
			args: args{digits: "2"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
