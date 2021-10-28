package tree

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				p: New([]int{1, 2, 3, 4, 5, 6, 7, null, 9}),
				q: New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			},
			want: false,
		},
		{
			args: args{
				p: New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
				q: New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*TreeNode, *TreeNode) bool{isSameTree, isSameTree2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.p, tt.args.q); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
