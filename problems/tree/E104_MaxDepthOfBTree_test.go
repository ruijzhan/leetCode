package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_maxDepth(t *testing.T) {
	fs := []func(*TreeNode) int{maxDepth, maxDepth1}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: New([]int{21, 7, 14, 1, 1, 2, 2, 3, 3}),
			},
			want: 4,
		},
		{
			args: args{
				root: New([]int{3, 9, 20, null, null, 15, 7}),
			},
			want: 3,
		},
		{
			args: args{
				root: New([]int{}),
			},
			want: 0,
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
