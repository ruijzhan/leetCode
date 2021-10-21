package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_isBalanced(t *testing.T) {
	fs := []func(*TreeNode) bool{isBalanced, isBalanced2}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: New([]int{3, 9, 20, null, null, 15, 7}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{1, 2, 2, 3, 3, null, null, 4, 4}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
