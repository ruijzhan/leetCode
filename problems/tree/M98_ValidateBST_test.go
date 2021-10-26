package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_isValidBST(t *testing.T) {
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
				root: New([]int{2, 1, 3}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{5, 1, 4, null, null, 3, 6}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{5, 4, 6, null, null, 3, 7}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{2, 2, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*TreeNode) bool{isValidBST, isValidBST2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.root.levelOrder(), got, tt.want)
				}
			})
		}
	}
}
