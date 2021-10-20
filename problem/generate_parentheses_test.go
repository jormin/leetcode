package problem

import (
	"reflect"
	"testing"
)

// TestDFS 测试深度优先搜索算法
func TestDFS(t *testing.T) {
	type args struct {
		path  string
		n     int
		left  int
		right int
		res   *[]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "01",
			args: args{
				path:  "",
				n:     2,
				left:  0,
				right: 0,
				res:   &[]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
			},
		)
	}
}

// TestGenerateParenthesis 测试括号生成
func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "01",
			args: args{n: 3},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "02",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "03",
			args: args{n: 0},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
