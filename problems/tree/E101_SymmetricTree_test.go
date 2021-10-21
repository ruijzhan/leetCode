package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_isSymmetric(t *testing.T) {
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
				root: New([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{1, 2, 2, null, 3, null, 3}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{1, 2, 2, 3, 4, 4, 2}),
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
	fs := []func(*TreeNode) bool{isSymmetric, isSymmetric2}
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
