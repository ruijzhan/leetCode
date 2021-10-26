package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root:      New([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}),
				targetSum: 22,
			},
			want: true,
		},
		{
			args: args{
				root:      New([]int{1, 2, 3}),
				targetSum: 5,
			},
			want: false,
		},
		{
			args: args{
				root:      New([]int{1, 2}),
				targetSum: 0,
			},
			want: false,
		},
		{
			args: args{
				root:      New([]int{1}),
				targetSum: 1,
			},
			want: true,
		},
		{
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
		{
			args: args{
				root:      New([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}),
				targetSum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*TreeNode, int) bool{hasPathSum} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root, tt.args.targetSum); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
