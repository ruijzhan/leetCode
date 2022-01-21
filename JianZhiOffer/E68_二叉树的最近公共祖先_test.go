package offer

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root []int
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 1},
			},
			want: &TreeNode{Val: 3},
		},
		{
			args: args{
				root: []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4},
				p:    &TreeNode{Val: 5},
				q:    &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor((*TreeNode)(tree.New(tt.args.root)), tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
