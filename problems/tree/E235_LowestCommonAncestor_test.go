package tree

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: New([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}),
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 8},
			},
			want: &TreeNode{Val: 6},
		},
		{
			args: args{
				root: New([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}),
				p:    &TreeNode{Val: 2},
				q:    &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
