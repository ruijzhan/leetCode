package tree

import (
	"reflect"
	"testing"
)

func Test_balanceBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: New([]int{1, null, 2, null, 3, null, 4, null, null}),
			},
			want: New([]int{2, 1, 3, null, null, null, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceBST(tt.args.root); !reflect.DeepEqual(got.inOrder(), tt.want.inOrder()) {
				t.Errorf("balanceBST() = %v, want %v", got.inOrder(), tt.want.inOrder())
			}
		})
	}
}
