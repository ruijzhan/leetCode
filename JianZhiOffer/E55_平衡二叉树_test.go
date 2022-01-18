package offer

import (
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: []int{3, 9, 20, null, null, 15, 7},
			},
			want: true,
		},
		{
			args: args{
				root: []int{1, 2, 2, 3, 3, null, null, 4, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced((*TreeNode)(tree.New(tt.args.root))); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
